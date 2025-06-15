package api

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func WriteStringJson(w http.ResponseWriter, statusCode int, payload string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(payload))
}
