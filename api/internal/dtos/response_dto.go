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

type StockDBDTO struct {
	Ticker     string  `json:"ticker"`
	Company    string  `json:"company"`
	Brokerage  string  `json:"brokerage"`
	Action     string  `json:"action"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	TargetFrom float64 `json:"target_from"`
	TargetTo   float64 `json:"target_to"`
}


