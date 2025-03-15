package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Blessed0314/tru-test/api/internal/dtos"
	"github.com/Blessed0314/tru-test/api/internal/models"
	"github.com/Blessed0314/tru-test/api/internal/utils"
	"github.com/Blessed0314/tru-test/api/pkg/db"
)

// ExternalAPIClient maneja las peticiones a la API externa
type ExternalAPIClient struct {
	BaseURL string
	Token   string
}

// FetchData obtiene los datos de la API externa
func (c *ExternalAPIClient) FetchData() ([]dtos.StockRatingDTO, error) {
	client := &http.Client{}
	var allItems []dtos.StockRatingDTO
	url := c.BaseURL

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("error creando request: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+c.Token)

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error enviando request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error de respuesta API: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error leyendo response: %v", err)
		}

		var response dtos.ApiData
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("error parseando JSON: %v", err)
		}

		allItems = append(allItems, response.Items...)

		if response.NextPage == "" {
			break
		}

		url = c.BaseURL + "?next_page=" + response.NextPage
	}

	return allItems, nil
}

// SaveData guarda los datos en la base de datos
func SaveData(stockData []dtos.StockRatingDTO) error {
	for _, stockRating := range stockData {
		log.Printf("📊 Procesando %s", stockRating.TargetFrom)
		var existingStockRating models.StockRating
		result := db.DB.Where("ticker = ?", stockRating.Ticker).First(&existingStockRating)

		// Convertir los precios correctamente
		targetFrom, errFrom := utils.ConvertPrice(stockRating.TargetFrom)
		if errFrom != nil {
			return fmt.Errorf("❌ Error convirtiendo TargetFrom (%s): %v", stockRating.TargetFrom, errFrom)
		}

		targetTo, errTo := utils.ConvertPrice(stockRating.TargetTo)
		if errTo != nil {
			return fmt.Errorf("❌ Error convirtiendo TargetTo (%s): %v", stockRating.TargetTo, errTo)
		}

		if result.Error != nil { // Si no existe, creamos un nuevo registro
			newStock := models.StockRating{
				Ticker:     stockRating.Ticker,
				Company:    stockRating.Company,
				Brokerage:  stockRating.Brokerage,
				Action:     stockRating.Action,
				RatingFrom: stockRating.RatingFrom,
				RatingTo:   stockRating.RatingTo,
				TargetFrom: targetFrom,
				TargetTo:   targetTo,
			}

			if err := db.DB.Create(&newStock).Error; err != nil {
				return fmt.Errorf("❌ Error guardando en BD: %v", err)
			}
		} else { // Si existe, lo actualizamos
			existingStockRating.Company = stockRating.Company
			existingStockRating.Brokerage = stockRating.Brokerage
			existingStockRating.Action = stockRating.Action
			existingStockRating.RatingFrom = stockRating.RatingFrom
			existingStockRating.RatingTo = stockRating.RatingTo
			existingStockRating.TargetFrom = targetFrom
			existingStockRating.TargetTo = targetTo

			if err := db.DB.Save(&existingStockRating).Error; err != nil {
				return fmt.Errorf("❌ Error actualizando BD: %v", err)
			}
		}
	}

	return nil
}

