package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type data struct {
	Files    map[string]string
	Filename string
	Content  string
}

func main() {
	files := map[string]string{
		"lunch_orders.txt":  "adamt: meat\nadamy: chicken\njames: rice",
		"new_hires.private": "EMPTY",
		"current_user":      "adamt"}

	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := data{files, "", ""}
		RenderTemplate(w, "index.html", data)
	})

	r.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		filename := r.PostFormValue("filename")

		if filename == "new_hires.private" {
			data, _ := ioutil.ReadFile(os.Args[2])
			w.Write(data)
			return
		}

		file, ok := files[filename]

		data := data{files, filename, ""}
		if !ok {
			data.Content = "File not found"
		} else {
			data.Content = string(file)
		}

		RenderTemplate(w, "index.html", data)
	})

	http.ListenAndServe(":80", r)
}
