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

type StockRecommendationDTO struct {
	Ticker   string  `json:"ticker"`
	Company  string  `json:"company"`
	Broker   string  `json:"brokerage"`
	Target   float64 `json:"target_change"`
	Rating   string  `json:"rating"`
	Score    float64 `json:"score"`
}

