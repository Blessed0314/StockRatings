package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Blessed0314/tru-test/api/internal/services"
	"github.com/Blessed0314/tru-test/api/internal/utils"
	"github.com/gorilla/mux"
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
        utils.SendResponse(w, r, http.StatusInternalServerError, "Error getting recommendations", nil)
        return
    }

    // Aplicar paginación antes de enviar la respuesta
    paginatedResponse := utils.Paginate(recommendations, page, pageSize)

    utils.SendResponse(w, r, http.StatusOK, "Recommendations OK", paginatedResponse)
}

func GetAllStocksHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := validatePaginationParams(r)

    stocks, err := services.GetAllStocks()
    if err != nil {
        utils.SendResponse(w, r, http.StatusInternalServerError, "Error getting stocks", nil)
        return
    }

    // Aplicar paginación antes de enviar la respuesta
    paginatedResponse := utils.Paginate(stocks, page, pageSize)

    utils.SendResponse(w, r, http.StatusOK, "Stocks OK", paginatedResponse)
}

func GetStockByTickerHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ticker := vars["ticker"]
    if ticker == "" {
        utils.SendResponse(w, r, http.StatusBadRequest, "Ticker parameter is required", nil)
        return
    }

    stock, err := services.GetStockByTicker(ticker)

    if errors.Is(err, services.ErrStockNotFound) {
        utils.SendResponse(w, r, http.StatusNotFound, "Stock not found", nil)
        return
    }

    if err != nil {
        utils.SendResponse(w, r, http.StatusInternalServerError, "Error getting stock", nil)
        return
    }

    utils.SendResponse(w, r, http.StatusOK, "Stock OK", stock)
}

func GetStocksByTickerPrefixHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ticker := vars["tickerPrefix"]
    if ticker == "" {
        utils.SendResponse(w, r, http.StatusBadRequest, "Ticker parameter is required", nil)
        return
    }

    page, pageSize := validatePaginationParams(r)

    stocks, err := services.GetStocksByTickerPrefix(ticker)
    if err != nil {
        utils.SendResponse(w, r, http.StatusInternalServerError, "Error getting stocks", nil)
        return
    }

    // Aplicar paginación antes de enviar la respuesta
    paginatedResponse := utils.Paginate(stocks, page, pageSize)

    utils.SendResponse(w, r, http.StatusOK, "Stocks OK", paginatedResponse)
}