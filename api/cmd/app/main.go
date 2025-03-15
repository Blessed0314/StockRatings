package main

import (
	"log"
	"net/http"

	"github.com/Blessed0314/tru-test/api/pkg/db"
	"github.com/Blessed0314/tru-test/api/internal/routes"
	"github.com/Blessed0314/tru-test/api/internal/models"
)

func main() {
	db.InitDB()
	err := db.DB.AutoMigrate(&models.StockRating{})
	if err != nil {
		log.Fatal("Error al migrar la base de datos", err)
	}

	mainRoutes := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":3001", mainRoutes))
}