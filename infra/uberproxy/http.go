package main

import (
	"io"
	"net/http"
)

func initUPRsp(rsp http.ResponseWriter) {
	rsp.Header().Add("Server", "UberProxy")
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			if k == "Server" {
				dst.Set(k, v)
			} else {
				dst.Add(k, v)
			}
		}
	}
}

func copyResponse(rw http.ResponseWriter, resp *http.Response) error {
	copyHeader(rw.Header(), resp.Header)
	rw.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()

	_, err := io.Copy(rw, resp.Body)
	return err
}
