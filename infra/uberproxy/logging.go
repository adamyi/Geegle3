package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
)

type uplogEntry struct {
	Request struct {
		Body   string `json:"body"`
		Header string `json:"headers"`
	} `json:"response"`
	Response struct {
		StatusCode int         `json:"statusCode"`
		Body       string      `json:"body"`
		Header     http.Header `json:"headers"`
	} `json:"response"`
}

func WrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		entry := uplogEntry{}
		rs, _ := httputil.DumpRequest(req, false)
		entry.Request.Header = string(rs)
		// http.MaxBytesReader might be better but let's just use io.LimitedReader since we are doing the wrapped logger.
		limitedReader := &io.LimitedReader{R: req.Body, N: 10485760}
		reqcontent, lrerr := ioutil.ReadAll(limitedReader)
		req.Body = ioutil.NopCloser(bytes.NewReader(reqcontent))
		entry.Request.Body = string(reqcontent)
		buf := new(bytes.Buffer)
		lrw := newUplogResponseWriter(buf, w, &entry)
		if lrerr != nil {
			returnError(UPError{Code: http.StatusBadRequest, Title: "You issued a malformed request", Description: "Entity Too Large"}, lrw)
		} else if limitedReader.N < 1 {
			returnError(UPError{Code: http.StatusRequestEntityTooLarge, Title: "You issued a malformed request", Description: "Entity Too Large"}, lrw)
		} else {
			wrappedHandler.ServeHTTP(lrw, req)
		}
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
