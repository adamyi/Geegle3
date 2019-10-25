package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func CGIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "healthz")
	w.Header().Add("Access-Control-Allow-Origin", "https://cli-relay.corp.geegle.org")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("OK"))
}

func check(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func main() {
	err := cgi.Serve(http.HandlerFunc(CGIHandler))
	check(err, "cannot serve request")
}
