package main

import (
	"math/rand"
	"net/http"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	initPrivateIP()
	initNetworkOptions()
	readConfig()

	go http.ListenAndServe(":80", http.HandlerFunc(redirectSSL))
	server := buildSSLServer()
	server.ListenAndServeTLS("", "")
}
