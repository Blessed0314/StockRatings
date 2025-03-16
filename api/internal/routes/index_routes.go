package routes

import (
	"github.com/Blessed0314/tru-test/api/internal/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	routes := mux.NewRouter()
	api := routes.PathPrefix("/api").Subrouter()
	api.HandleFunc("/data", handlers.GetDataHandler).Methods("GET")

	stock := routes.PathPrefix("/stock").Subrouter()
	stock.HandleFunc("/recommendations", handlers.GetStockRecommendationsHandler).Methods("GET")

	return routes
}