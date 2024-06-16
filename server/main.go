package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./server/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Print("Starting server at port 3003")
	if err := http.ListenAndServe("127.0.0.1:3003", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Form Handler: method:%v", req.Method)

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err %v", err)
		return
	}
	log.Printf("POST request successful")

	name := req.FormValue("name")
	address := req.FormValue("address")

	log.Printf("Name(%v), Address(%v)", name, address)

	http.Redirect(w, req, "/form.html", http.StatusFound)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Hello Handler: method:%v", req.Method)

	if req.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}
