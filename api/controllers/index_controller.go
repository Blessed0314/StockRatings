package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Blessed0314/tru-test/api/db"
	"github.com/Blessed0314/tru-test/api/models"
	"github.com/Blessed0314/tru-test/api/utils"
	"github.com/joho/godotenv"
)

func GetData(w http.ResponseWriter, r *http.Request) {
    err := godotenv.Load()
    if err != nil {
        utils.SendResponse(w, http.StatusInternalServerError, "Error loading .env file")
        return
    }

    token := os.Getenv("API_TOKEN")
    if token == "" {
        utils.SendResponse(w, http.StatusInternalServerError, "API_TOKEN is not set")
        return
    }
    
    baseURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
    url := baseURL

    client := &http.Client{}

    for {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            utils.SendResponse(w, http.StatusInternalServerError, "Error creating request")
            return
        }

        req.Header.Set("Authorization", "Bearer "+ token)

        resp, err := client.Do(req)
        if err != nil {
            utils.SendResponse(w, http.StatusInternalServerError, "Error sending request")
            return
        }
        defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
            utils.SendResponse(w, resp.StatusCode, fmt.Sprintf("Error response from API: %s", resp.Status))
            return
        }

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            utils.SendResponse(w, http.StatusInternalServerError, "Error reading response")
            return
        }

        // Estructura para deserializar la respuesta JSON
        var response struct {
            Items    []models.StockRating `json:"items"`
            NextPage string               `json:"next_page"`
        }

        // Deserializar la respuesta en la estructura
        err = json.Unmarshal(body, &response)
        if err != nil {
            utils.SendResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling response: %v", err))
            return
        }

        // Guardar los datos en la base de datos
        for _, stockRating := range response.Items {
            var existingStockRating models.StockRating
            result := db.DB.Where("ticker = ?", stockRating.Ticker).First(&existingStockRating)
            if result.Error != nil {
                // Si no se encuentra el registro, crear uno nuevo
                result = db.DB.Create(&stockRating)
                if result.Error != nil {
                    utils.SendResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error saving data to database: %v", result.Error))
                    return
                }
            } else {
                // Si se encuentra el registro, actualizar los campos necesarios
                existingStockRating.Company = stockRating.Company
                existingStockRating.Brokerage = stockRating.Brokerage
                existingStockRating.Action = stockRating.Action
                existingStockRating.RatingFrom = stockRating.RatingFrom
                existingStockRating.RatingTo = stockRating.RatingTo
                existingStockRating.TargetFrom = stockRating.TargetFrom
                existingStockRating.TargetTo = stockRating.TargetTo
                db.DB.Save(&existingStockRating)
            }
        }

        // Si no hay más páginas, salir del bucle
        if response.NextPage == "" {
            break
        }

        // Actualizar la URL para la siguiente página
        url = baseURL + "?next_page=" + response.NextPage
    }

    utils.SendResponse(w, http.StatusOK, "Data saved successfully")
}