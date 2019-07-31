package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type data struct {
	URL     string
	Content string
}

func main() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "index.html", data{})
	})

	r.Methods("GET").Path("/private").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You must be authenticated and attached to the internal VPN to view this page."))
	})

	r.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.PostFormValue("search")

		if url == "http://localhost/private" {
			flag, _ := ioutil.ReadFile("flag")
			RenderTemplate(w, "index.html", data{url, string(flag)})
			return
		}

		client := http.Client{Timeout: 5 * time.Second}
		out, err := client.Get(url)

		var pagecontent string
		if err != nil {
			pagecontent = err.Error()
		} else {
			defer out.Body.Close()
			page, err := ioutil.ReadAll(out.Body)

			if err != nil {
				pagecontent = err.Error()
			} else {
				pagecontent = string(page)
			}
		}

		RenderTemplate(w, "index.html", data{url, pagecontent})
	})

	http.ListenAndServe(":8003", r)
}
