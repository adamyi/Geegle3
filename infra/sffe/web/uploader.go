package web

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/adamyi/geegle-sffe/context"
)

type Flag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ApiStoreRequest struct {
	FileName string `json:"filename"`
	Content  []byte `json:"content"`
	Flags    []Flag `json:"flags"`
}

func StoreFile(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	tknStr := r.Header.Get("X-Geegle-JWT")
	service, err := getJwtServiceName(tknStr, []byte("superSecretJWTKEY"))
	if err != nil {
		writeJSONError(w, "Invalid JWT", http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeJSONBackendError(w, err)
		return
	}

	var req ApiStoreRequest
	err = json.Unmarshal(b, &req)

	if err != nil {
		writeJSONError(w, "JSON Error", http.StatusBadRequest)
		return
	}

	sort.SliceStable(req.Flags, func(i, j int) bool {
		return req.Flags[i].Name < req.Flags[j].Name
	})

	var urlb strings.Builder
	fmt.Fprintf(&urlb, "%x/", md5.Sum([]byte(service)))
	for _, flag := range req.Flags {
		fmt.Fprintf(&urlb, "%s=%s/", flag.Name, flag.Value)
	}

	urlb.WriteString(req.FileName)

	url := urlb.String()

	f := context.File{
		Content: req.Content,
		Time:    time.Now(),
	}

	err = ctx.Put(url, &f)

	if err != nil {
		writeJSONBackendError(w, err)
		return
	}

	writeJSONUrl(w, url)
}
