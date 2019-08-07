package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func redirectSSL(rsp http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(rsp, req, target,
		http.StatusTemporaryRedirect)
}

func buildSSLServer() http.Server {
	cfg := &tls.Config{}
	var err error
	cfg.Certificates = make([]tls.Certificate, 3)
	cfg.Certificates[0], err = tls.LoadX509KeyPair("certs/geegle.org.pem", "certs/geegle.org.key")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates[1], err = tls.LoadX509KeyPair("certs/corp.geegle.org.pem", "certs/corp.geegle.org.key")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates[2], err = tls.LoadX509KeyPair("certs/apps.geegle.org.pem", "certs/apps.geegle.org.key")
	if err != nil {
		log.Fatal(err)
	}

	cfg.BuildNameToCertificate()

	server := http.Server{
		Addr:      ":443",
		Handler:   http.HandlerFunc(handleUP),
		TLSConfig: cfg,
	}
	return server
}
