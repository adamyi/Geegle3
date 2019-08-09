package web

import (
	"fmt"
	"net/http"

	"geegle.org/infra/sffe/context"
)

func initSffeRsp(w http.ResponseWriter) {
	w.Header().Add("Server", "sffe")
}

func ListenAndServe(addr string, ctx *context.Context) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/store/", func(w http.ResponseWriter, r *http.Request) {
		initSffeRsp(w)
		StoreFile(ctx, w, r)
	})
	mux.HandleFunc("/s/", func(w http.ResponseWriter, r *http.Request) {
		initSffeRsp(w)
		ServeFile(ctx, w, r)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		initSffeRsp(w)
		fmt.Fprintln(w, "👍")
	})

	return http.ListenAndServe(addr, mux)
}
