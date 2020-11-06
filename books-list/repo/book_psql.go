package bookrepo

import (
	"books-list/model"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooksRepo(db *sql.DB, book model.Book, res []model.Book) ([]model.Book, error) {
	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return []model.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		res = append(res, book)
	}

	if err != nil {
		return []model.Book{}, err
	}

	return res, nil
}
