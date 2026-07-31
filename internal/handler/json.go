package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	encoder.Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {

	w.WriteHeader(status)

	WriteJSON(w, map[string]string{
		"error": message,
	})
}
