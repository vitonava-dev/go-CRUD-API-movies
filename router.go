package main

import "github.com/gorilla/mux"

func initRouter(r *mux.Router) {

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

}
