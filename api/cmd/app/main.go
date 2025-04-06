package main

import (
	"log"
	"net/http"

	"github.com/Blessed0314/tru-test/api/internal/models"
	"github.com/Blessed0314/tru-test/api/internal/routes"
	"github.com/Blessed0314/tru-test/api/internal/middlewares"
	"github.com/Blessed0314/tru-test/api/pkg/db"
)

func main() {
	db.InitDB()
	err := db.DB.AutoMigrate(&models.StockRating{})
	if err != nil {
		log.Fatal("Error when migrating the database", err)
	}

	mainRoutes := routes.InitRouter()
	corsRoutes := middlewares.ConfigureCORS(mainRoutes)

	log.Fatal(http.ListenAndServe(":3001", corsRoutes))
}