package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

  log.Print("Starting server at port 8082")
  if err := http.ListenAndServe(":8082", nil); err != nil {
    log.Fatal(err)
  }
}

func formHandler(w http.ResponseWriter, req *http.Request) {
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
}
