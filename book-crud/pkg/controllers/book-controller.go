package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/book-crud/pkg/models"
	"example.com/book-crud/pkg/utils"
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

	fmt.Println("Found book:", book.ID, book.Name)

	w.Header().Set("Content-Type", "application/json")

	var obj any = book

	if book.ID == 0 {
		obj = struct{}{}
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	res, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func CreateBook(w http.ResponseWriter, req *http.Request) {
	book := &models.Book{}
	utils.ParseBody(req.Body, &book)

	b := book.CreateBook()

	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	book := models.DeleteBook(uint(id))

	res, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	update := &models.Book{}
	utils.ParseBody(req.Body, update)

	book, db := models.GetBookById(uint(id))

	if update.Name != "" {
		book.Name = update.Name
	}
	if update.Author != "" {
		book.Author = update.Author
	}
	if update.Publication != "" {
		book.Publication = update.Publication
	}

	db.Save(book)

	res, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
