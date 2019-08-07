package web

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getJwtServiceName(tknStr string, JwtKey []byte) (string, error) {
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

	if !strings.HasSuffix(claims.Username, "@services.geegle.org") {
		return "", fmt.Errorf("Not a service account")
	}

	return claims.Username[:len(claims.Username)-20], nil
}
