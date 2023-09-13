package routes

import (
	"net/http"
)

func Setup() {

	http.HandleFunc("/api/v1/download", nil)

}
