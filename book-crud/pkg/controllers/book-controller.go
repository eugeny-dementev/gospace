package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/book-crud/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, req *http.Request) {
	books := models.GetAllBooks()

	res, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
