package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		perm, _ := r.Cookie("permission")
		if perm != nil {
			RenderTemplate(w, "index.html", perm.Value)
		} else {
			RenderTemplate(w, "index.html", "")
		}

	})

	r.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expiration := time.Now().Add(1 * time.Minute)
		cookie := http.Cookie{Name: "permission", Value: "intern", Expires: expiration}
		http.SetCookie(w, &cookie)

		RenderTemplate(w, "index.html", cookie.Value)
	})

	http.ListenAndServe(":80", r)
}
