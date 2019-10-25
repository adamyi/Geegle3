package web

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getJwtUsername(tknStr string, JwtKey []byte) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", fmt.Errorf("JWT Invalid")
	}

	return claims.Username, nil
}

func getJwtServiceName(tknStr string, JwtKey []byte) (string, error) {
	username, err := getJwtUsername(tknStr, JwtKey)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(username, "@services.geegle.org") {
		return "", fmt.Errorf("Not a service account")
	}

	return username[:len(username)-20], nil
}

func getJwtLDAPName(tknStr string, JwtKey []byte) (string, error) {
	username, err := getJwtUsername(tknStr, JwtKey)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(username, "@geegle.org") {
		return "", fmt.Errorf("Not a corp account")
	}

	return username[:len(username)-11], nil
}
