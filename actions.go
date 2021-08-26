package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running on http://localhost:8080")
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin limites", 2013, "Desconocido"},
		Movie{"Batman Begins", 2005, "Nolan"},
		Movie{"A todo gas", 2005, "Juan Antonio"},
	}
	json.NewEncoder(w).Encode(movies)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie ID %s", movie_id)
}