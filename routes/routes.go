package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheusferreira165/tablescraper/controllers"
	"github.com/rs/cors"
)

func Setup() *mux.Router {

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	r.Use(c.Handler)

	r.HandleFunc("/api/v1/download", controllers.GenerateTokenDownload).Methods("POST")
	r.HandleFunc("/api/v1/download/{token}", controllers.DownloadTable).Methods("GET")

	return r
}
