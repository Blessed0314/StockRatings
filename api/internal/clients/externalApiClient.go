package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Blessed0314/tru-test/api/internal/dtos"
)

// ExternalAPIClient maneja la comunicación con la API externa
type ExternalAPIClient struct {
	BaseURL string
	Token   string
}

// FetchData obtiene los datos desde la API externa
func (c *ExternalAPIClient) FetchData() ([]dtos.StockRatingDTO, error) {
	client := &http.Client{}
	var allItems []dtos.StockRatingDTO
	url := c.BaseURL

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("error creando la petición: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+c.Token)

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error enviando la petición: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error en la respuesta de la API: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error leyendo la respuesta: %v", err)
		}

		var response dtos.ApiData
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("error deserializando la respuesta: %v", err)
		}

		allItems = append(allItems, response.Items...)

		if response.NextPage == "" {
			break
		}

		url = c.BaseURL + "?next_page=" + response.NextPage
	}

	return allItems, nil
}