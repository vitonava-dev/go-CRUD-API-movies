package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range movies {
		if item.ID == params["id"] {

			item.Status = false
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "204 Movie not found", http.StatusNoContent)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "204 Movie not found", http.StatusNoContent)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movie.Status = true
	movies = append(movies, &movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			fmt.Println(movie)
			item.Isbn = movie.Isbn
			item.Director = movie.Director
			item.Title = movie.Title

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "204 Movie not found", http.StatusNoContent)

}
