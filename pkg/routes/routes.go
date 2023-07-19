package routes

import (
	"github.com/Alhiane/goapi-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	// ...
	router := mux.NewRouter()

	router.HandleFunc("/api/books/", controllers.GetNotes).Methods("GET")
	router.HandleFunc("/api/books/", controllers.GetNotes).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.GetNote).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetNote).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.GetNote).Methods("DELETE")

	return router
}
