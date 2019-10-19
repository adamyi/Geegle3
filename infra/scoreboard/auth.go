package web

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getJwtUserName(tknStr string, JwtKey []byte) (string, error) {
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

	if claims.Service != "geemail-backend@services.geegle.org" {
		return "", fmt.Errorf("JWT not for geemail")
	}

	return claims.Username, nil
}

func confirmFromGeemail(tknStr string) error {
	_, err := getJwtServiceName(tknStr, []byte("superSecretJWTKEY"))
	if err != nil {
		return err
	}

	return nil

}
