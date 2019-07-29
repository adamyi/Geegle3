package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	files := map[string]string{
		"lunch_orders.txt":  "adamt: meat",
		"new_hires.private": "<URL>",
		"current_user":      "geegle"}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "index.html", files)
	})

	r.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		filename := r.PostFormValue("filename")

		file, ok := files[filename]

		if !ok {
			w.Write([]byte("File not found"))
		}

		w.Write([]byte(file))
	})

	http.ListenAndServe(":8001", r)
}
