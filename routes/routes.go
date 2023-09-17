package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusferreira165/tablescraper/controllers"
	"github.com/rs/cors"
)

func Setup() *mux.Router {

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	r.HandleFunc("/api/v1/download/", controllers.DownloadTable).Methods("GET")
	r.HandleFunc("/api/v1/download/{token}", controllers.DownloadTable).Methods("POST")
	http.Handle("/", c.Handler(http.DefaultServeMux))

	return r
}
