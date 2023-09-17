package main

import (
	"net/http"

	"github.com/matheusferreira165/tablescraper/routes"
)

func main() {
	m := routes.Setup()
	http.ListenAndServe(":3000", m)
}
