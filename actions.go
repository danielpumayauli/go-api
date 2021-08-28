package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func response(w http.ResponseWriter, status int, results interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running on http://localhost:8080")
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := movieColletion.Find(nil).Sort("-_id").All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}
	response(w, 200, results)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	results := Movie{}
	err := movieColletion.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	response(w, 200, results)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = movieColletion.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		// w.Write([]byte("500 - Something bad happened!"))
		return
	}
	response(w, 200, movie_data)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)

	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = movieColletion.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	response(w, 200, movie_data)
}

func RemoveMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	err := movieColletion.RemoveId((oid))

	if err != nil {
		w.WriteHeader(404)
		return
	}
	result := new(Message)
	result.setStatus("success")
	result.setMessage("The movie with ID " + movie_id + " was successfully removed")

	response(w, 200, result)
}
