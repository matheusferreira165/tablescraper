package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func DownloadTable(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	token := vars["token"]

	fileName, exists := downloadTokens[token]
	if !exists {
		http.Error(w, "Token inválido", http.StatusBadRequest)
		return
	}

	tmpFile, err := os.Open(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=TableData.csv")
	w.Header().Set("Content-Type", "text/csv")

	_, err = io.Copy(w, tmpFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	os.Remove(fileName)
}
