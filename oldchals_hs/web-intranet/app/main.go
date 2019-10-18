package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type data struct {
	URL     string
	Content string
}

func main() {
	r := mux.NewRouter()

	r.Host("internal.docs").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "http://internal.docs/robots\nhttp://internal.docs/interns\nhttp://internal.docs/employees")
	})

	r.Host("internal.docs").Path("/robots").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "internal.docs/*")
	})

	r.Host("internal.docs").Path("/interns").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "No intern docs currently... (maybe an intern project for next year?)")
	})

	r.Host("internal.docs").Path("/employees").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Employee DOCS can't be accessed via the VPN, they must be accessed via internal machines")
	})

	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "index.html", data{})
	})

	r.Methods("GET").Path("/private").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You must be authenticated and attached to the internal VPN to view this page."))
	})

	r.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.PostFormValue("search")

		if url == "http://localhost/private" || url == "whatevertheurlis" {
			flag, _ := ioutil.ReadFile(os.Args[2])
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

	http.ListenAndServe(":80", r)
}
