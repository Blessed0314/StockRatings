package routes

import (
	"github.com/Blessed0314/tru-test/api/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	routes := mux.NewRouter()
	api := routes.PathPrefix("/api").Subrouter()
	
	api.HandleFunc("/data", controllers.GetData).Methods("GET")

	return routes
}