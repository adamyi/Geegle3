package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handleCLIRelay(rsp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/cli-relay.py" {
		rsp.Header().Set("Content-Type", "text/x-python")
		body, _ := ioutil.ReadFile(os.Args[3] + "/cli-relay.py")
		fmt.Fprint(rsp, string(body))
	} else {
		body, _ := ioutil.ReadFile(os.Args[3] + "/ws.html")
		fmt.Fprint(rsp, string(body))
	}
}
