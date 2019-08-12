package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Configuration struct {
	RedisHost     string
	ListenAddress string
	JwtKey        []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var _redis redis.Conn

func redisConnect() {
	var err error
	_redis, err = redis.Dial("tcp", _configuration.RedisHost)
	if err != nil {
		panic(err)
	}
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

func initMssRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "mss")
}

func handleMSS(rsp http.ResponseWriter, req *http.Request) {
	initMssRsp(rsp)

	tknStr := req.Header.Get("X-Geegle-JWT")

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return _configuration.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Signature Invalid")
			rsp.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("JWT Error")
		log.Println(err.Error())
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		log.Println("JWT Invalid")
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !strings.HasSuffix(claims.Username, "@services.geegle.org") {
		log.Println(claims.Username)
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}

	service := claims.Username[:len(claims.Username)-20]
	key := service + req.URL.Path

	switch req.Method {
	case "GET":
		value, err := redis.String(_redis.Do("GET", key))
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(rsp, value)
		return
	case "POST":
		value, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = _redis.Do("SET", key, value)
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
		} else {
			expires := req.Header.Get("X-MSS-Expires")
			if expires != "" {
				_, err = _redis.Do("EXPIRE", key, expires)
				if err != nil {
					rsp.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
		io.WriteString(rsp, "OK")
	default:
		http.Error(rsp, "Method not allowed.", http.StatusMethodNotAllowed)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	time.Sleep(5 * time.Second)
	redisConnect()
	http.HandleFunc("/", handleMSS)
	err := http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
