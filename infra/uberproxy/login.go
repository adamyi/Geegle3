package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

func verifyPassword(email, password string) bool {

	data, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{email, password})
	if err != nil {
		fmt.Println(err)
		return false
	}

	r, err := http.Post("http://gaia.corp.geegle.org/api/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return r.StatusCode == 200
}

func handleLogin(rsp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tmpl := template.Must(template.ParseFiles(os.Args[3] + "/login.html"))
		tmpl.Execute(rsp, "")
	} else if req.Method == "POST" {
		req.ParseForm()
		username := strings.ToLower(req.Form.Get("username"))
		if !strings.HasSuffix(username, "@geegle.org") {
			username = username + "@geegle.org"
		}
		password := req.Form.Get("password")

		if !verifyPassword(username, password) {
			tmpl := template.Must(template.ParseFiles(os.Args[3] + "/login.html"))
			tmpl.Execute(rsp, "Incorrect password")
			return
		}

		expirationTime := time.Now().Add(24 * 30 * time.Hour)
		pclaims := Claims{
			Username: username,
			Service:  "uberproxy@services.geegle.org",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		ptoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pclaims)
		ptstr, err := ptoken.SignedString(_configuration.SignKey)
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
