package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"geegle.org/infra/sffe/context"
)

func StoreFile(ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	tknStr := r.Header.Get("X-Geegle-JWT")
	service, err := getJwtServiceName(tknStr, []byte("superSecretJWTKEY"))
	if err != nil {
		writeJSONError(w, "Invalid JWT", http.StatusUnauthorized)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeJSONBackendError(w, err)
		return
	}

	var req context.StoreRequest
	err = json.Unmarshal(b, &req)

	if err != nil {
		writeJSONError(w, "JSON Error", http.StatusBadRequest)
		return
	}

	req.Service = service

	url, err := context.DoStoreFile(ctx, &req)

	if err != nil {
		writeJSONBackendError(w, err)
		return
	}

	writeJSONUrl(w, url)
}
