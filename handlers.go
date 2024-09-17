package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getBooks(w http.ResponseWriter, r *http.Request){
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()


	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			http.Error(w, "Error scanning book data", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// Prépare l'instruction SQL pour insertion
	sqlStatement := `INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id`
	err = db.QueryRow(sqlStatement, newBook.Title, newBook.Author).Scan(&newBook.ID)

	if err != nil {
		http.Error(w, "Unable to execute the query", http.StatusInternalServerError)
		log.Println("Error inserting book:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook) // Retourne le nouveau livre avec son ID
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	//Extract book id from URL
	idstr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	//Prepar SQL Statement
	sqlStatement := `SELECT id, title, author FROM books WHERE id = $1`
	var book Book
	err = db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		http.Error(w, "Unable to execute the query", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/books/delete/")
	id , err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	//Prepar SQL statement
	sqlStatement := `DELETE FROM books WHERE id = $1 RETURNING id`
	var book Book
	err = db.QueryRow(sqlStatement, id).Scan(&book.ID)
	if err != nil {
		http.Error(w, "Invalid query", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid Book", http.StatusBadRequest)
		return
	}
	idstr := strings.TrimPrefix(r.URL.Path, "/books/update/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	sqlStatement := `UPDATE books SET title = $2, author = $3 WHERE id = $1 RETURNING id, title, author`
	err = db.QueryRow(sqlStatement, id, updatedBook.Title, updatedBook.Author).Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author)
	if err != nil {
		http.Error(w, "Invalid query", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content_type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}