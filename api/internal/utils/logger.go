package utils

import (
	"log"
	"net/http"
	"time"
)

func CreateLog(r *http.Request, statusCode int, message string) {

	path := r.URL.Path
	log.Printf("[%s] %s %s - code: %d %s", time.Now().Format(time.RFC3339), r.Method, path, statusCode, message)
}
