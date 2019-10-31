package main

import (
	"crypto/rsa"
	"encoding/json"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
	JwtPubKey     []byte
	SignKey       *rsa.PrivateKey
	VerifyKey     *rsa.PublicKey
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
	_configuration.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(_configuration.JwtKey)
	if err != nil {
		panic(err)
	}
	_configuration.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(_configuration.JwtPubKey)
	if err != nil {
		panic(err)
	}

}
