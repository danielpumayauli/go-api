package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/movies", Movies)
	router.HandleFunc("/movies/{id}", ShowMovie)

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running on http://localhost:8080")
}

func Movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Movies Page")
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Movie ID %s", movie_id)
}