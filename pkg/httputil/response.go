package httputil

import (
	"encoding/json"
	"net/http"
	"time"
)

func writeJsonResponse[T any](w http.ResponseWriter, statusCode int, data Response[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func Success[T any](w http.ResponseWriter, statusCode int, data T, message string) {
	r := Response[T]{
		Data:      data,
		Success:   true,
		Message:   message,
		Timestamp: time.Now().UTC().Unix(),
	}
	writeJsonResponse(w, statusCode, r)
}

func Fail[T any](w http.ResponseWriter, statusCode int, message string) {
	r := Response[T]{
		Success:   false,
		Message:   message,
		Timestamp: time.Now().UTC().Unix(),
	}
	writeJsonResponse(w, statusCode, r)
}
