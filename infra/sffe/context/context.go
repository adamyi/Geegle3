package context

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	filesDbFilename = "files.db"
)

// File is the value part of a shortcut.
type File struct {
	Content []byte    `json:"content"`
	Time    time.Time `json:"time"`
}

type Flag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StoreRequest struct {
	FileName string `json:"filename"`
	Content  []byte `json:"content"`
	Flags    []Flag `json:"flags"`
	Service  string `json:"service"`
}

// Serialize this File into the given Writer.
func (o *File) write(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, o.Time.UnixNano()); err != nil {
		return err
	}

	if _, err := w.Write([]byte(o.Content)); err != nil {
		return err
	}

	return nil
}

// Deserialize this File from the given Reader.
func (o *File) read(r io.Reader) error {
	var t int64
	if err := binary.Read(r, binary.LittleEndian, &t); err != nil {
		return err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	o.Content = b
	o.Time = time.Unix(0, t)
	return nil
}

// Context provides access to the data store.
type Context struct {
	path string
	db   *leveldb.DB
	lck  sync.Mutex
	id   uint64
}

// Open the context using path as the data store location.
func Open(path string) (*Context, error) {
	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	// open the database
	db, err := leveldb.OpenFile(filepath.Join(path, filesDbFilename), nil)
	if err != nil {
		return nil, err
	}

	return &Context{
		path: path,
		db:   db,
	}, nil
}

// Close the resources associated with this context.
func (c *Context) Close() error {
	return c.db.Close()
}

// Get retreives a file from the data store.
func (c *Context) Get(name string) (*File, error) {
	val, err := c.db.Get([]byte(name), nil)
	if err != nil {
		return nil, err
	}

	rt := &File{}
	if err := rt.read(bytes.NewBuffer(val)); err != nil {
		return nil, err
	}

	return rt, nil
}

// Put stores a new file in the data store.
func (c *Context) Put(key string, rt *File) error {
	var buf bytes.Buffer
	if err := rt.write(&buf); err != nil {
		return err
	}

	return c.db.Put([]byte(key), buf.Bytes(), &opt.WriteOptions{Sync: true})
}

// Del removes an existing file from the data store.
func (c *Context) Del(key string) error {
	return c.db.Delete([]byte(key), &opt.WriteOptions{Sync: true})
}

// GetAll gets everything in the db to dump it out for backup purposes
func (c *Context) GetAll() (map[string]File, error) {
	files := map[string]File{}
	iter := c.db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		rt := &File{}
		if err := rt.read(bytes.NewBuffer(val)); err != nil {
			return nil, err
		}
		files[string(key[:])] = *rt
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return files, nil
}

func DoStoreFile(ctx *Context, req *StoreRequest) (string, error) {

	sort.SliceStable(req.Flags, func(i, j int) bool {
		return req.Flags[i].Name < req.Flags[j].Name
	})

	var urlb strings.Builder
	fmt.Fprintf(&urlb, "%x/", md5.Sum([]byte(req.Service)))
	for _, flag := range req.Flags {
		fmt.Fprintf(&urlb, "%s=%s/", flag.Name, flag.Value)
	}

	urlb.WriteString(req.FileName)

	url := urlb.String()

	f := File{
		Content: req.Content,
		Time:    time.Now(),
	}

	log.Println("Storing at: " + url)

	return url, ctx.Put(url, &f)
}

func InitFiles(ctx *Context, filepath string) {
	filepath += "/"
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	var sr []StoreRequest
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			jsonFile, err := os.Open(filepath + f.Name())
			if err != nil {
				log.Println(err)
				continue
			}
			jsonContent, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				log.Println(err)
				continue
			}
			err = json.Unmarshal(jsonContent, &sr)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, r := range sr {
				storedFile, err := os.Open(filepath + r.FileName)
				if err != nil {
					log.Println(err)
					continue
				}
				r.Content, err = ioutil.ReadAll(storedFile)
				if err != nil {
					log.Println(err)
					continue
				}
				DoStoreFile(ctx, &r)
			}
		}
	}
}
