package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(os.Args[1])))

	log.Fatal(http.ListenAndServe(":80", nil))
}
