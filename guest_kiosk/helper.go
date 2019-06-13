package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/oxtoacart/bpool"
)

var bufpool *bpool.BufferPool

var isProd bool

func Setup() {
	flag.BoolVar(&isProd, "production", false, "if true, we start HTTPS server")
	flag.Parse()
	bufpool = bpool.NewBufferPool(64)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("templates/layouts/base.html", "templates/"+name))

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func makeServer(handler http.Handler) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}
}
func SetupHTTPS(r http.Handler) {
	if isProd {
		dataDir := "."

		httpsSrv := makeServer(r)
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("guest.geegle.org"),
			Cache:      autocert.DirCache(dataDir),
		}
		httpsSrv.Addr = ":443"
		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			// serve HTTP, which will redirect automatically to HTTPS
			h := m.HTTPHandler(nil)
			log.Fatal(http.ListenAndServe(":http", h))
		}()

		// serve HTTPS!
		log.Fatal(httpsSrv.ListenAndServeTLS("", ""))

	} else { // Devonly

		http.ListenAndServe(":80", r)
	}
}
