package utils

import (
	"encoding/json"
	"net/http"
	"time"

    "github.com/Blessed0314/tru-test/api/internal/dtos"
)



func SendResponse(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := dtos.Response{
        StatusCode: statusCode,
        Message:    message,
        Timestamp:  time.Now(),
    }
    json.NewEncoder(w).Encode(response)
}