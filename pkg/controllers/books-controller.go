package controllers

import (
	"fmt"
	"net/http"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	// ...

	fmt.Fprintf(w, "GetNotes Endpoint Hit")

	return
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	// ...

	return
}
