package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var _db *sql.DB

func main() {
	rand.Seed(time.Now().UnixNano())
	initPrivateIP()
	initNetworkOptions()
	readConfig()

	_db, err = sql.Open(_configuration.DbType, _configuration.DbAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()

	go http.ListenAndServe(":80", http.HandlerFunc(redirectSSL))
	server := buildSSLServer()
	server.ListenAndServeTLS("", "")
}
