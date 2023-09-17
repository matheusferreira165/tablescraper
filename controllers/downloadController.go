package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/matheusferreira165/tablescraper/models"
	"github.com/matheusferreira165/tablescraper/services"
)

func DownloadTable(w http.ResponseWriter, r *http.Request) {
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

	tmpFile, err := os.Open(filecsv.Name())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=TableData.csv")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "text/csv")

	_, err = io.Copy(w, tmpFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
