package main

import (
	"log"
	"net/http"

	"example.com/book-crud/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("127.0.0.1:3007", router))
}
