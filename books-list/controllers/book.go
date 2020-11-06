package controllers

import (
	"books-list/driver"
	"books-list/model"
	bookrepo "books-list/repo"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Controller struct{}

var res []model.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create variable to store the current book in the next/scan loop
		var book model.Book

		//create response slice to return to the front end
		res = []model.Book{}

		repo := bookrepo.BookRepository{}

		res, err := repo.GetBooksRepo(db, book, res)
		driver.LogFatal(err)

		//encode our res slice and send it back to the front end
		json.NewEncoder(w).Encode(res)
	}
}
