package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matheusferreira165/tablescraper/models"
	"github.com/matheusferreira165/tablescraper/services"
)

var downloadTokens = make(map[string]string)

func GenerateToken(w http.ResponseWriter, r *http.Request) {

	token := services.TokenGenerator()
	downloadTokens[token] = "TableData.csv"

	downloadURL := fmt.Sprintf("/api/v1/download/%s", token)

	response := models.DownloadLink{
		Token: token,
		URL:   downloadURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
