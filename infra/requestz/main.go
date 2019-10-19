package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var _configuration = Configuration{}

func readConfig() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func initRZRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "requestz")
	rsp.Header().Add("Content-Type", "text/plain")
}

func handleRZ(rsp http.ResponseWriter, req *http.Request) {
	initRZRsp(rsp)

	rs, _ := httputil.DumpRequest(req, true)

	rsp.Write(rs)

	tknStr := req.Header.Get("X-Geegle-JWT")

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return _configuration.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rsp.Write([]byte("JWT: signature invalid\n"))
			return
		}
		rsp.Write([]byte("JWT: error\n"))
		rsp.Write([]byte(err.Error()))
		return
	}

	if !tkn.Valid {
		rsp.Write([]byte("JWT: invalid\n"))
		return
	}

	s, _ := json.Marshal(claims)

	rsp.Write(s)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	http.HandleFunc("/", handleRZ)
	err := http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
