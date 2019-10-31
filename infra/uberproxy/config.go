package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	ListenAddress string
	JwtKey        []byte
	DbType        string
	DbAddress     string
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

}
