package handlers

import (
    "net/http"
    "strconv"

    "github.com/Blessed0314/tru-test/api/internal/services"
    "github.com/Blessed0314/tru-test/api/internal/utils"
)


func validatePaginationParams(r *http.Request) (int, int) {
    page, err := strconv.Atoi(r.URL.Query().Get("page"))
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
    if err != nil || pageSize < 1 {
        pageSize = 10
    }

    return page, pageSize
}

func GetStockRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := validatePaginationParams(r)

    recommendations, err := services.GetStockRecommendations()
    if err != nil {
        utils.SendResponse(w, http.StatusInternalServerError, "Error getting recommendations", nil)
        return
    }

    // Aplicar paginación antes de enviar la respuesta
    paginatedResponse := utils.Paginate(recommendations, page, pageSize)

    utils.SendResponse(w, http.StatusOK, "Recommendations OK", paginatedResponse)
}

func GetAllStocksHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := validatePaginationParams(r)

    stocks, err := services.GetAllStocks()
    if err != nil {
        utils.SendResponse(w, http.StatusInternalServerError, "Error getting stocks", nil)
        return
    }

    // Aplicar paginación antes de enviar la respuesta
    paginatedResponse := utils.Paginate(stocks, page, pageSize)

    utils.SendResponse(w, http.StatusOK, "Stocks OK", paginatedResponse)
}