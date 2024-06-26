package routes

import (
	"example.com/book-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
  router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
  router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
  router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
  router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
  router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
