package utils

import (
	"encoding/json"
	"net/http"
	"time"

    "github.com/Blessed0314/tru-test/api/internal/dtos"
)



func SendResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string, data any) {
    CreateLog(r, statusCode, message)
    
    response := dtos.Response{
        StatusCode: statusCode,
        Message:    message,
        Timestamp:  time.Now(),
    }

    if data != nil {
        response.Data = data
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}