package dtos

import (
	"time"
)

type Response struct {
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	Data	any `json:"data,omitempty"`
}