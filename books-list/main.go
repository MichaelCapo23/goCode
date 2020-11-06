package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Server running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	params := mux.Vars(r)

	rows := db.QueryRow("SELECT * FROM books WHERE id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(&book)

}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id;", &book.Title, &book.Author, &book.Year, &book.ID)
	driver.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("DELETE FROM books WHERE id=$1", params["id"])

	rowsDeleted, err := result.RowsAffected()
	driver.LogFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
