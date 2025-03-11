package main

import (
	"log"
	"net/http"

	"github.com/Blessed0314/tru-test/api/routes"
)

func main() {
	InitDB()

	mainRoutes := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":3001", mainRoutes))
}