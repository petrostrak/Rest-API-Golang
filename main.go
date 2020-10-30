package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Every func that is a route handler has to have (w http.ResponseWriter, r *http.Request) parameters.
// They are similar to (req, res)

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get book by id
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "332251", Title: "First Book", Author: &Author{Firstname: "Petros", Lastname: "Trak"}})
	books = append(books, Book{ID: "2", Isbn: "433562", Title: "Second Book", Author: &Author{Firstname: "Eleni", Lastname: "Apost"}})
	books = append(books, Book{ID: "3", Isbn: "985746", Title: "Third Book", Author: &Author{Firstname: "Greg", Lastname: "Bail"}})

	// Route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
