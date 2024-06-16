package main

import (
	"fmt"
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
  if req.URL.Path != "/hello" {
    http.Error(w, "404 Not Found", http.StatusNotFound)
    return
  }

  if req.Method != "GET" {
    http.Error(w, "Method not supported", http.StatusNotFound)
    return;
  }

  fmt.Fprintf(w, "Hello")
}
