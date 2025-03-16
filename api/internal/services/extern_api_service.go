package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/Blessed0314/tru-test/api/internal/dtos"
    "github.com/Blessed0314/tru-test/api/internal/models"
    "github.com/Blessed0314/tru-test/api/internal/repository"
    "github.com/Blessed0314/tru-test/api/internal/utils"
)

// ExternalAPIClient handles requests to the external API
type ExternalAPIClient struct {
    BaseURL string
    Token   string
}

// FetchData retrieves data from the external API
func (c *ExternalAPIClient) FetchData() ([]dtos.StockRatingDTO, error) {
    client := &http.Client{}
    var allItems []dtos.StockRatingDTO
    url := c.BaseURL

    for {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return nil, fmt.Errorf("error creating request: %v", err)
        }

        req.Header.Set("Authorization", "Bearer "+c.Token)

        resp, err := client.Do(req)
        if err != nil {
            return nil, fmt.Errorf("error sending request: %v", err)
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("API response error: %s", resp.Status)
        }

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, fmt.Errorf("error reading response: %v", err)
        }

        var response dtos.ApiData
        err = json.Unmarshal(body, &response)
        if err != nil {
            return nil, fmt.Errorf("error parsing JSON: %v", err)
        }

        allItems = append(allItems, response.Items...)

        if response.NextPage == "" {
            break
        }

        url = c.BaseURL + "?next_page=" + response.NextPage
    }

    return allItems, nil
}

func SaveData(stockData []dtos.StockRatingDTO) error {
    stockRepo := repository.NewStockRepository()

    for _, stockRating := range stockData {
        var existingStockRating models.StockRating
        result := stockRepo.DB.Where("ticker = ?", stockRating.Ticker).First(&existingStockRating)

        // Convert prices correctly
        targetFrom, errFrom := utils.ConvertPrice(stockRating.TargetFrom)
        if errFrom != nil {
            return fmt.Errorf("❌ Error converting TargetFrom (%s): %v", stockRating.TargetFrom, errFrom)
        }

        targetTo, errTo := utils.ConvertPrice(stockRating.TargetTo)
        if errTo != nil {
            return fmt.Errorf("❌ Error converting TargetTo (%s): %v", stockRating.TargetTo, errTo)
        }

        if result.Error != nil { // If it doesn't exist, create a new record
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

            if err := stockRepo.Save(&newStock); err != nil {
                return fmt.Errorf("❌ Error saving to DB: %v", err)
            }
        } else { // If it exists, update it
            existingStockRating.Company = stockRating.Company
            existingStockRating.Brokerage = stockRating.Brokerage
            existingStockRating.Action = stockRating.Action
            existingStockRating.RatingFrom = stockRating.RatingFrom
            existingStockRating.RatingTo = stockRating.RatingTo
            existingStockRating.TargetFrom = targetFrom
            existingStockRating.TargetTo = targetTo

            if err := stockRepo.Update(&existingStockRating); err != nil {
                return fmt.Errorf("❌ Error updating DB: %v", err)
            }
        }
    }

    return nil
}