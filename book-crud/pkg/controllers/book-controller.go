package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/book-crud/pkg/models"
	"github.com/gorilla/mux"
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

func GetBookById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
  if err != nil {
    log.Fatal(err)
  }

	book, _ := models.GetBookById(uint(id))

  res, err := json.Marshal(book)
  if err != nil {
    log.Fatal(err)
  }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
