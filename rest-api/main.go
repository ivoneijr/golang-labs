package main

import (
	"encoding/json"
	"net/http"

	// "encoding/json"
	"log"
	"math/rand"
	"strconv"

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
	FirsName string `json:firstname`
	LastName string `json:lastname`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock
	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&book)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			book.ID = params["id"]
			books = append(books, book)
			return
		}
	}

	json.NewEncoder(w).Encode(books)
}

func delBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement db
	books = append(books, Book{ID: "1", Isbn: "123easf", Title: "Ivo Book", Author: &Author{FirsName: "Ivo", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "222easf", Title: "Anne Book", Author: &Author{FirsName: "Ari", LastName: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "333easf", Title: "Valen Book", Author: &Author{FirsName: "Val", LastName: "Doe"}})
	books = append(books, Book{ID: "4", Isbn: "333easf", Title: "Amarula Book", Author: &Author{FirsName: "Lulux", LastName: "Doe"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", delBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
