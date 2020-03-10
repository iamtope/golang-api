package main // define package main for every main go file

import (
	"encoding/json"  // for handing json 
	"log"           // for logging
	"net/http"     // to manage http requests
	"math/rand"    // to generate random numbers
	"strconv"		// string converter
	"github.com/gorilla/mux" // Your imported package
)
// Book Structs (Our Model)
type Book struct {
	ID string `json: "id"`
	isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

// Author struct 
type Author struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}
// Initialize book as a slice book struct

var books[]Book


// create our route handler function
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get any params
	// loop through the books
	for _, item:= range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
	
}

// create new books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	 var book Book
	 _ = json.NewDecoder(r.Body).Decode(&book)
	 book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock Id which is not safe as it could generate an id twice. 
	 books = append(books, book)
	 json.NewEncoder(w).Encode(book)

	
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item:= range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books) 

}


func main(){  // define a function main with the func keyword
	// initiate the mux router
	r := mux.NewRouter()  // := is a type inference declaring of a variable in go

	// Mock data instead of implementing a database
	books = append(books, Book{ID: "1", isbn: "495930", Title: "Book One", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", isbn: "495958", Title: "Book Two", Author: &Author {Firstname: "Isaac", Lastname: "Newton"}})

	// Create our route handlers or endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET") 
	r.HandleFunc("/api/books/{id}", getBook ).Methods("GET") 
	r.HandleFunc("/api/books", createBook).Methods("POST") 
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT") 
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE") 
	
// create our server to listen on port 8000
log.Fatal(http.ListenAndServe(":8000", r)) 

}

