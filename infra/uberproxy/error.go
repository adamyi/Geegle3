package main

import (
	"html/template"
	"net/http"
	"os"
)

type UPError struct {
	Title       string
	Description string
	Code        int
}

func returnError(err UPError, rsp http.ResponseWriter) {
	rsp.WriteHeader(err.Code)
	tmpl := template.Must(template.ParseFiles(os.Args[3] + "/error.html"))
	tmpl.Execute(rsp, err)
}
