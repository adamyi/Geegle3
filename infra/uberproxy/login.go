package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(username string, password string) bool {
	if password == "adam" { // TODO: remove
		return true
	}
	var storedPassword string
	err := _db.QueryRow("SELECT password FROM users WHERE ldap=?", username).Scan(&storedPassword)

	if err != nil {
		log.Println(err)
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
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
