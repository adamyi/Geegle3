package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

type uplogEntry struct {
	Request  string `json:"request"`
	Response struct {
		StatusCode int         `json:"statusCode"`
		Body       string      `json:"body"`
		Header     http.Header `json:"headers"`
	} `json:"response"`
}

func WrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		entry := uplogEntry{}

		rs, _ := httputil.DumpRequest(req, true)
		entry.Request = string(rs)
		buf := new(bytes.Buffer)

		lrw := newUplogResponseWriter(buf, w, &entry)
		wrappedHandler.ServeHTTP(lrw, req)

		entry.Response.Body = buf.String()
		entry.Response.Header = lrw.Header()
		fmt.Println(lrw.Header())
		estr, _ := json.Marshal(entry)
		fmt.Println(string(estr))
	})
}

type uplogResponseWriter struct {
	file  io.Writer
	resp  http.ResponseWriter
	multi io.Writer
	entry *uplogEntry
}

func newUplogResponseWriter(file io.Writer, resp http.ResponseWriter, entry *uplogEntry) http.ResponseWriter {
	multi := io.MultiWriter(file, resp)
	return &uplogResponseWriter{
		file:  file,
		resp:  resp,
		multi: multi,
		entry: entry,
	}
}

// implement http.ResponseWriter
// https://golang.org/pkg/net/http/#ResponseWriter
func (w *uplogResponseWriter) Header() http.Header {
	return w.resp.Header()
}

func (w *uplogResponseWriter) Write(b []byte) (int, error) {
	return w.multi.Write(b)
}

func (w *uplogResponseWriter) WriteHeader(i int) {
	w.entry.Response.StatusCode = i
	w.resp.WriteHeader(i)
}

// websocket needs this
func (w *uplogResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := w.resp.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, errors.New("Error in hijacker")
}
