package main

import (
	"encoding/json"
	"fmt"      /* used for printing to the console */
	"log"      /* used for logging errors */
	"net/http" /* used for creating the HTTP server */ /* used for encoding and decoding JSON */ /* used for generating random IDs */ /* used for converting strings to integers and vice versa */

	"github.com/gorilla/mux" /* used for routing HTTP requests */
)

/* Movie struct represents a movie with its ID, ISBN, title, and director */
type Movie struct {
	ID       string    `json:"id"`       /* unique identifier for the movie */
	Isbn     string    `json:"isbn"`     /* International Standard Book Number for the movie */
	Title    string    `json:"title"`    /* title of the movie */
	Director *Director `json:"director"` /* pointer to a Director struct representing the director of the movie */
}

/* Director struct represents a director with their first and last name */
type Director struct {
	Firstname string `json:"firstname"` /* first name of the director */
	Lastname  string `json:"lastname"`  /* last name of the director */
}

var movies []Movie /* slice to hold the list of movies */

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") /* set the response header to indicate that the content type is JSON */
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter() /* create a new router using the Gorilla Mux package */

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})    /* add a sample movie to the movies slice */
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}}) /* add another sample movie to the movies slice */

	r.HandleFunc("/movies", getMovies).Methods("GET")           /* route for getting all movies */
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       /* route for getting a single movie by ID */
	r.HandleFunc("/movies", createMovie).Methods("POST")        /* route for creating a new movie */
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    /* route for updating an existing movie */
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") /* route for deleting a movie by ID */

	fmt.Printf("Starting server at port 8000\n") /* print a message indicating that the server is starting */
	log.Fatal(http.ListenAndServe(":8000", r))   /* start the HTTP server on port 8000 and log any errors */

}
