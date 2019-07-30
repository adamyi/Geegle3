package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
}

type Claims struct {
	Username string `json:"username"`
	Service  string `json:"service"`
	jwt.StandardClaims
}

var _configuration = Configuration{}

func readConfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&_configuration)
	if err != nil {
		panic(err)
	}
}

func initUPRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "UberProxy")
}

func handleUP(rsp http.ResponseWriter, req *http.Request) {
	initUPRsp(rsp)
	if req.Host == "login.corp.geegle.org" {
		handleLogin(rsp, req)
		return
	}
	c, err := req.Cookie("uberproxy_auth")
	full_url := req.Host + req.RequestURI
	if err != nil {
		http.Redirect(rsp, req, "http://login.corp.geegle.org/?return_url="+url.QueryEscape("http://"+full_url), 303)
		return
	}
	tknStr := c.Value
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

	if claims.Service != "uberproxy@services.geegle.org" {
		log.Println(claims.Service)
		rsp.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: check if user has permission to access this site

	expirationTime := time.Now().Add(5 * time.Minute)
	pclaims := Claims{
		Username: claims.Username,
		Service:  claims.Service, //TODO: check service name
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	ptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pclaims)
	ptstr, err := ptoken.SignedString(_configuration.JwtKey)
	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}

	preq, err := http.NewRequest(req.Method, "http://"+full_url, req.Body)
	if err != nil {
		fmt.Println("test")
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	for name, value := range req.Header {
		//TODO: remove uberproxy auth cookie from forwarded request
		preq.Header.Set(name, value[0])
	}
	preq.Header.Set("X-Geegle-JWT", ptstr)
	client := &http.Client{}
	presp, err := client.Do(preq)
	if err != nil {
		fmt.Println("haha")
		fmt.Println(err.Error())
		rsp.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Body.Close()
	for k, v := range presp.Header {
		rsp.Header().Set(k, v[0])
	}
	rsp.WriteHeader(presp.StatusCode)
	io.Copy(rsp, presp.Body)
	presp.Body.Close()
}

func handleLogin(rsp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//TODO: cache this
		body, _ := ioutil.ReadFile("templates/login.html")
		fmt.Fprint(rsp, string(body))
	} else if req.Method == "POST" {
		req.ParseForm()
		username := strings.ToLower(req.Form.Get("username"))
		if !strings.HasSuffix(username, "@geegle.org") {
			username = username + "@geegle.org"
		}
		_ = req.Form.Get("password")
		//TODO: verify password
		expirationTime := time.Now().Add(1 * time.Hour)
		pclaims := Claims{
			Username: username,
			Service:  "uberproxy@services.geegle.org",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		ptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pclaims)
		ptstr, err := ptoken.SignedString(_configuration.JwtKey)
		if err != nil {
			rsp.WriteHeader(http.StatusInternalServerError)
			return
		}
		authcookie := &http.Cookie{Name: "uberproxy_auth", Value: ptstr, HttpOnly: true, Domain: "corp.geegle.org"}
		http.SetCookie(rsp, authcookie)
		http.Redirect(rsp, req, req.URL.Query()["return_url"][0], 303)
	} else {
		rsp.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	readConfig()
	http.HandleFunc("/", handleUP)
	err := http.ListenAndServe(_configuration.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
