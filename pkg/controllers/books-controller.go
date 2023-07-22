package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alhiane/goapi-crud/pkg/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// ...
	allBooks, err := models.GetBooks()
	if err != nil {
		// ...
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allBooks)
	w.Write([]byte("GetBooks"))

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// ...

	return
}
