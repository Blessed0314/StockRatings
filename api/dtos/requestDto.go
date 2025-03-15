package dtos

import "github.com/Blessed0314/tru-test/api/models"

type ApiData = struct {
	Items    []models.StockRating `json:"items"`
	NextPage string              `json:"next_page"`
}
