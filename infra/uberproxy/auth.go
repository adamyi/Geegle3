package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Service  string `json:"service"`
	jwt.StandardClaims
}

func getUsername(req *http.Request) string {
	c, err := req.Cookie("uberproxy_auth")
	if err != nil {
		return "anonymous@services.geegle.org"
	}
	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return _configuration.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Signature Invalid")
			return "anonymous@services.geegle.org"
		}
		log.Println("JWT Error")
		log.Println(err.Error())
		return "anonymous@services.geegle.org"
	}

	if !tkn.Valid {
		log.Println("JWT Invalid")
		return "anonymous@services.geegle.org"
	}

	if claims.Service != "uberproxy@services.geegle.org" {
		return "anonymous@services.geegle.org"
	}

	return claims.Username
}
