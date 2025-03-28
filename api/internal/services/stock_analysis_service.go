package services

import (
	"math"
	"sort"
	"time"

	"github.com/Blessed0314/tru-test/api/internal/dtos"
	"github.com/Blessed0314/tru-test/api/internal/models"
	"github.com/Blessed0314/tru-test/api/internal/repository"
	"github.com/Blessed0314/tru-test/api/internal/utils"
)

// Calcula el puntaje de la acción
func calculateScore(stock models.StockRating) float64 {
    // Si el precio objetivo bajó, descartamos (-∞)
    if stock.TargetTo < stock.TargetFrom {
        return -math.MaxFloat64
    }

    // Variación porcentual en el precio objetivo
    percentageChange := ((stock.TargetTo - stock.TargetFrom) / stock.TargetFrom) * 100

    // Bonus por mejora en la calificación
    ratingBoost := 0.0
    if stock.RatingFrom != stock.RatingTo {
        ratingBoost = 5.0
    }

    // Penalización por antigüedad (datos más nuevos son mejores)
    daysOld := time.Since(stock.Time).Hours() / 24
    recencyBoost := math.Max(0, 30-daysOld) / 30 * 10

    // Puntuación final
    return math.Round((percentageChange + ratingBoost + recencyBoost) * 100) / 100
}

// Obtiene las mejores acciones para invertir
func GetStockRecommendations() ([]dtos.StockRecommendationDTO, error) {
    stockRepo := repository.NewStockRepository()
    stocks, err := stockRepo.GetAll()
    if err != nil {
        return nil, err
    }

    var recommendations []dtos.StockRecommendationDTO
    for _, stock := range stocks {
        score := calculateScore(stock)
        if score == -math.MaxFloat64 {
            continue // Ignorar acciones con Target decreciente
        }

        recommendations = append(recommendations, dtos.StockRecommendationDTO{
            Ticker:  stock.Ticker,
            Company: stock.Company,
            Broker:  stock.Brokerage,
            Target:  utils.ConvertToString(stock.TargetTo - stock.TargetFrom),
            Rating:  stock.RatingTo,
            Score:   score,
        })
    }

    // Ordenar por mejor puntuación
    sort.Slice(recommendations, func(i, j int) bool {
        return recommendations[i].Score > recommendations[j].Score
    })

    return recommendations, nil
}