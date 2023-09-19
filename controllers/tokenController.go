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

	var message models.Messages

	data, err := io.ReadAll(r.Body)
	if err != nil {
		message.Message = "error reading the website body"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	var download models.TableLink
	if err := json.Unmarshal(data, &download); err != nil {
		message.Message = "JSON deserialization error: " + err.Error()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	filecsv, err := services.GenerateCsv(download.Link)
	if err != nil {
		message.Message = "error generate csv archive"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
