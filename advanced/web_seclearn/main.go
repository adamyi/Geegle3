package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"sort"
	// "net/url"
	"os"
	// "os/exec"
	"strconv"
	"strings"
	// "strings"
	"time"

	"code.sajari.com/word2vec"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
}

var _configuration = Configuration{}
var _db *sql.DB
var _model *word2vec.Model

var validWord = regexp.MustCompile(`^[a-z ]+$`).MatchString

func readConfig() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func initSlRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "seclearn")
}

// TODO: XSRF
func newWord(rsp http.ResponseWriter, req *http.Request) {
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	tknStr := req.Header.Get("X-Geegle-JWT")
	user, err := getJwtUserName(tknStr, _configuration.JwtKey)
	if err != nil {
		fmt.Printf(err.Error())
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	word := strings.ToLower(req.URL.Query().Get("word"))
	if !validWord(word) {
		rsp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rsp, "Can only contain letters and space!")
		return
	}
	var id int
	err = _db.QueryRow("select max(id) from words where user = ?", user).Scan(&id)
	if err == nil {
		id += 1
	} else {
		id = 0
	}
	_, err = _db.Exec("insert into words (user, id, word) values(?, ?, ?)", user, id, word)
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(rsp, req, "/details/"+strconv.Itoa(id)+"?"+req.URL.RawQuery, http.StatusTemporaryRedirect)

}

func checkWord(rsp http.ResponseWriter, req *http.Request) {
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}

	tknStr := req.Header.Get("X-Geegle-JWT")
	user, err := getJwtUserName(tknStr, _configuration.JwtKey)
	if err != nil {
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}
	var inter struct {
		Word    string
		History []string
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}
	err = _db.QueryRow("select word from words where user = ? and id = ?", user, id).Scan(&(inter.Word))
	if err != nil {
		rsp.WriteHeader(http.StatusNotFound)
		return
	}
	inter.History = append(req.URL.Query()["h"], inter.Word)
	sort.Strings(inter.History)
	//inter.EscapedWord = url.QueryEscape(inter.Word)
	RenderTemplate(rsp, "word.html", inter)
}

func similarWords(rsp http.ResponseWriter, req *http.Request) {
	// time.Sleep(1200 * time.Millisecond)
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}
	rsp.Header().Set("Content-Type", "application/javascript")
	vars := mux.Vars(req)
	word := vars["word"]
	// give them hope, pretend this is a cool JSONP vuln
	cb := req.URL.Query().Get("cb")
	if cb != "callback" && cb != "callback2" {
		// it's actually not
		fmt.Fprintln(rsp, "console.log(\"bad hacker\")")
		return
	}
	// out, err := exec.Command("python", "word2vec.py", word).Output()
	words := strings.Fields(word)
	coe := 1.0 / float32(len(words))
	expr := word2vec.Expr{}
	for _, word := range words {
		expr.Add(coe, word)
	}
	matches, err := _model.CosN(expr, 10)
	var out []byte
	if err != nil {
		out = []byte("[]")
	} else {
		out, err = json.Marshal(matches)
		if err != nil {
			out = []byte("[]")
		}
	}
	fmt.Fprintf(rsp, "%s(%s)", cb, strings.TrimSuffix(string(out), "\n"))
}

func index(rsp http.ResponseWriter, req *http.Request) {
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}
	var inter struct{}
	RenderTemplate(rsp, "index.html", inter)
}

func ad(rsp http.ResponseWriter, req *http.Request) {
	time.Sleep(1200 * time.Millisecond)
	initSlRsp(rsp)
	if req.Method == "OPTIONS" {
		return
	}
	rsp.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintln(rsp, "document.getElementById('ad').innerHTML = 'We are now dogfooding our internal document typesetting service! Give it a go at <a href=\"https://docs.corp.geegle.org\">docs</a>!'")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	var err error
	_db, err = sql.Open("sqlite3", os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()
	modelfile, err := os.Open(os.Args[4])
	if err != nil {
		log.Fatalf("error loading model: %v", err)
	}
	_model, err = word2vec.FromReader(modelfile)
	if err != nil {
		log.Fatalf("error loading model: %v", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/details/{id:[0-9]+}", checkWord)
	r.HandleFunc("/api/addconcept", newWord)
	r.HandleFunc("/api/recommendation/{word}/jsonp", similarWords)
	r.HandleFunc("/api/wait_for_ad", ad)
	http.Handle("/", r)
	err = http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
