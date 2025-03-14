package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
    StatusCode int       `json:"status_code"`
    Message    string    `json:"message"`
    Timestamp  time.Time `json:"timestamp"`
}

func SendResponse(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := Response{
        StatusCode: statusCode,
        Message:    message,
        Timestamp:  time.Now(),
    }
    json.NewEncoder(w).Encode(response)
}