package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/b1da4a0e-fbf7-11e9-aad5-362b9e155667/", http.StripPrefix("/b1da4a0e-fbf7-11e9-aad5-362b9e155667/", http.FileServer(http.Dir("/src/"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
