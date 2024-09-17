package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	initDB()
	defer db.Close()

	//routes
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here is my Api ! Welcome !")
	})

	http.HandleFunc("/books", getBooks) //GET all books
	http.HandleFunc("/books/", getBookByID) //GET one book
	http.HandleFunc("/books/create", createBook) //POST a new book
	http.HandleFunc("/books/delete/", deleteBook) // DELETE a book
	http.HandleFunc("/books/update/", updateBook) //UPDATE a book

	// Start server
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}