package main

import (
	"net/http"

	"github.com/Alhiane/goapi-crud/pkg/routes"
)

func main() {
	routes.SetupRoutes()

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}

}
