package models

import (
	"gorm.io/gorm"
)

type StockRating struct {
	gorm.Model
	ID         uint    `json:"id" gorm:"primaryKey"`
	Ticker     string  `json:"ticker" gorm:"size:10;not null; unique"`
	Company    string  `json:"company" gorm:"size:255;not null"`
	Brokerage  string  `json:"brokerage" gorm:"size:255;not null"`
	Action     string  `json:"action" gorm:"size:100;not null"`
	RatingFrom string  `json:"rating_from" gorm:"size:100"`
	RatingTo   string  `json:"rating_to" gorm:"size:100"`
	TargetFrom string `json:"target_from" gorm:"size:100"`
    TargetTo   string `json:"target_to" gorm:"size:100"`
}