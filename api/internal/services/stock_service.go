package services

import (
	"github.com/Blessed0314/tru-test/api/internal/dtos"
	"github.com/Blessed0314/tru-test/api/internal/models"
	"github.com/Blessed0314/tru-test/api/internal/repository"
)

func GetAllStocks() ([]dtos.StockDBDTO, error) {
    stockRepo := repository.NewStockRepository()
    stockRatings, err := stockRepo.GetAll()
    if err != nil {
        return nil, err
    }

    var stockDBDTOs []dtos.StockDBDTO
    for _, stockRating := range stockRatings {
        stockDBDTO := mapStockRatingToStockDBDTO(stockRating)
        stockDBDTOs = append(stockDBDTOs, stockDBDTO)
    }

    return stockDBDTOs, nil
}

func mapStockRatingToStockDBDTO(stockRating models.StockRating) dtos.StockDBDTO {
    return dtos.StockDBDTO{
        Ticker:     stockRating.Ticker,
        Company:    stockRating.Company,
        Brokerage:  stockRating.Brokerage,
        Action:     stockRating.Action,
        RatingFrom: stockRating.RatingFrom,
        RatingTo:   stockRating.RatingTo,
        TargetFrom: stockRating.TargetFrom,
        TargetTo:   stockRating.TargetTo,
    }
}