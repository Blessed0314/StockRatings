package main

import (
	"log"
	"net/http"

	"github.com/Blessed0314/tru-test/api/db"
	"github.com/Blessed0314/tru-test/api/routes"
	"github.com/Blessed0314/tru-test/api/models"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.StockRating{})

	mainRoutes := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":3001", mainRoutes))
}