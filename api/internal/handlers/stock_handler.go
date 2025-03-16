package handlers

import (
    "net/http"
    "os"

    "github.com/Blessed0314/tru-test/api/internal/services"
    "github.com/Blessed0314/tru-test/api/internal/utils"
)

// GetDataHandler handles the HTTP request to get data
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
    token := os.Getenv("API_TOKEN")
    if token == "" {
        utils.SendResponse(w, http.StatusInternalServerError, "API_TOKEN is not configured", nil)
        return
    }

    client := &services.ExternalAPIClient{
        BaseURL: "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list",
        Token:   token,
    }

    data, err := client.FetchData()
    if err != nil {
        utils.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    err = services.SaveData(data)
    if err != nil {
        utils.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    utils.SendResponse(w, http.StatusOK, "Data saved successfully", nil)
}

func GetStockRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
    recommendations, err := services.GetStockRecommendations()
    if err != nil {
        utils.SendResponse(
            w, http.StatusInternalServerError, "Error getting recommendations", nil)
        return
    }

    utils.SendResponse(w, http.StatusOK, "Recommendations OK", recommendations)
}