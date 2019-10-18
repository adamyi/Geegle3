package main

import (
	"html/template"
	"net/http"
	"os"
)

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(os.Args[1]+"/layouts/base.html", os.Args[1]+"/"+name))

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
