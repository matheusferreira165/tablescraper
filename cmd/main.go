package main

import (
	"fmt"
	"net/http"

	"github.com/matheusferreira165/tablescraper/routes"
)

func main() {
	m := routes.Setup()
	fmt.Println("SERVIDOR INICIADO COM SUCESSO")
	http.ListenAndServe(":3000", m)
}
