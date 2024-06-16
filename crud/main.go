package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Director *Director `json:"director"`
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Year     int64     `json:"year"`
}

type Director struct {
	First string `json:"name"`
	Last  string `json:"last"`
}

func main() {
	router := mux.NewRouter()

	movies := make([]Movie, 0)
	movies = append(movies, Movie{
		&Director{
			"John",
			"Nolan",
		},
		"someid",
		"very specivic isbn",
		"Rookie",
		2024,
	})

	router.HandleFunc("/movies", func(w http.ResponseWriter, req *http.Request) {
		jsonBuffer, err := json.Marshal(movies)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonBuffer)
	}).Methods("GET")

	router.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := "someid"
		var movie Movie
		for _, m := range movies {
			if m.ID == id {
				movie = m
			}
		}
		jsonBuffer, err := json.Marshal(movie)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonBuffer)
	}).Methods("GET")

	router.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		// movie := Movie{}
		// json.NewDecoder(r.read).Decode(movie)
	}).Methods("POST")
	router.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("PUT")
	router.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("DELETE")

	fmt.Println("Server started on localhost:3005")
	http.ListenAndServe("127.0.0.1:3005", router)
}
