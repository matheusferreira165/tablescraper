package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matheusferreira165/tablescraper/models"
	"github.com/matheusferreira165/tablescraper/services"
)

var downloadTokens = make(map[string]string)

func GenerateTokenDownload(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var download models.TableLink
	if err := json.Unmarshal(data, &download); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filecsv, err := services.GenerateCsv(download.Link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := services.TokenGenerator()
	downloadTokens[token] = filecsv.Name()

	downloadURL := fmt.Sprintf("/api/v1/download/%s", token)

	response := models.DownloadLink{
		Token: token,
		URL:   downloadURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
