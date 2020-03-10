package main // define package main for every main go file

import (
	// "encoding/json"  // for handing json 
	"log"           // for logging
	"net/http"     // to manage http requests
	// "math/rand"    // to generate random numbers
	// "strconv"		// string converter
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
// create our route handler function
func getBooks(w http.ResponseWriter, r *http.Request) {
	
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// create new books
func createBook(w http.ResponseWriter, r *http.Request) {
	
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
}


func main(){  // define a function main with the func keyword
	// initiate the mux router
	r := mux.NewRouter()  // := is a type inference declaring of a variable in go

	// Create our route handlers or endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET") 
	r.HandleFunc("/api/books/{id}", getBook ).Methods("GET") 
	r.HandleFunc("/api/books", createBook).Methods("POST") 
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT") 
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE") 
	
// create our server to listen on port 8000
log.Fatal(http.ListenAndServe(":8000", r)) 

}

