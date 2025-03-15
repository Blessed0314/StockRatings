package handlers

import (
	"net/http"
	"os"

	"github.com/Blessed0314/tru-test/api/internal/services"
	"github.com/Blessed0314/tru-test/api/internal/utils"
)

// GetData maneja la petición HTTP para obtener datos
func GetData(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		utils.SendResponse(w, http.StatusInternalServerError, "API_TOKEN no está configurado")
		return
	}

	client := &services.ExternalAPIClient{
		BaseURL: "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list",
		Token:   token,
	}

	data, err := client.FetchData()
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = services.SaveData(data)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(w, http.StatusOK, "Datos guardados correctamente")
}
