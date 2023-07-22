package routes

import (
	"net/http"

	"github.com/Alhiane/goapi-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	// ...
	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/api/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/", controllers.GetBooks).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("DELETE")

	return router
}
