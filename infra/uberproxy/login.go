package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func handleLogin(rsp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//TODO: cache this
		body, _ := ioutil.ReadFile(os.Args[3] + "/login.html")
		fmt.Fprint(rsp, string(body))
	} else if req.Method == "POST" {
		req.ParseForm()
		username := strings.ToLower(req.Form.Get("username"))
		if !strings.HasSuffix(username, "@geegle.org") {
			username = username + "@geegle.org"
		}
		_ = req.Form.Get("password")
		//TODO: verify password
		expirationTime := time.Now().Add(24 * 30 * time.Hour)
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
			returnError(UPError{Code: http.StatusInternalServerError, Title: "Internal Server Error", Description: "Something went wrong while generating JWT"}, rsp)
			return
		}
		authcookie := &http.Cookie{Name: "uberproxy_auth", Value: ptstr, HttpOnly: true, Domain: "geegle.org"}
		http.SetCookie(rsp, authcookie)
		http.Redirect(rsp, req, req.URL.Query()["return_url"][0], http.StatusFound)
	} else {
		returnError(UPError{Code: http.StatusMethodNotAllowed, Title: "Method Not Allowed", Description: "Only GET and POST are supported"}, rsp)
	}
}
