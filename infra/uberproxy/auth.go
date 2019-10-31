package main

import (
	"errors"
	"log"
	"net/http"
        "regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Service  string `json:"service"`
	jwt.StandardClaims
}

var SubAccValid = regexp.MustCompile(`^[a-zA-Z\-_]+$`).MatchString

func getUsername(req *http.Request) (string, error) {
	username := getMainUsername(req)
	subacc := req.Header.Get("X-Geegle-SubAcc")
        if !(SubAccValid(subacc) && len(subacc) < 10) {
		return "", errors.New("invalid subacc")
        }
	if subacc != "" {
		s := strings.Split(username, "@")
		username = s[0] + "+" + subacc + "@" + s[1]
	}
	return username, nil
}

func getMainUsername(req *http.Request) string {
	c, err := req.Cookie("uberproxy_auth")
	var tknStr string
	if err != nil {
		if tknStr = req.Header.Get("X-Geegle-JWT"); tknStr == "" {
			sn, err := getServiceNameFromIP(strings.Split(req.RemoteAddr, ":")[0])
			if err != nil {
				return "anonymous@services.geegle.org"
			}
			return sn
		}
	} else {
		tknStr = c.Value
	}
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
