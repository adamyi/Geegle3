package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// The response type for all API responses.
type msg struct {
	Ok bool `json:"ok"`
}

type msgErr struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type msgUrl struct {
	Ok  bool   `json:"ok"`
	Url string `json:"url"`
}

// Encode the given data to JSON and send it to the client.
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err)
	}
}

// Encode a simple success msg and send it to the client.
func writeJSONOk(w http.ResponseWriter) {
	writeJSON(w, &msg{
		Ok: true,
	}, http.StatusOK)
}

// Encode an error response and send it to the client.
func writeJSONError(w http.ResponseWriter, err string, status int) {
	writeJSON(w, &msgErr{
		Ok:    false,
		Error: err,
	}, status)
}

// Encode a generic backend error and send it to the client.
func writeJSONBackendError(w http.ResponseWriter, err error) {
	log.Printf("[error] %s", err)
	writeJSONError(w, "backend error", http.StatusInternalServerError)
}

// Send given url to the client.
func writeJSONUrl(w http.ResponseWriter, url string) {
	writeJSON(w, &msgUrl{
		Ok:  true,
		Url: url,
	}, http.StatusOK)
}
