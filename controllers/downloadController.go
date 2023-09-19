package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matheusferreira165/tablescraper/models"
)

func DownloadTable(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	token := vars["token"]

	var message models.Messages

	fileName, exists := downloadTokens[token]
	if !exists {
		message.Message = "invalid token"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message)
		return
	}

	tmpFile, err := os.Open(fileName)
	if err != nil {
		message.Message = "failed to read file"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message)
		return
	}
	defer tmpFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, err = io.Copy(w, tmpFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(downloadTokens, token)
	os.Remove(fileName)
}
