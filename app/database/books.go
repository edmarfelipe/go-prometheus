package database

import (
	"database/sql"

	"github.com/edmarfelipe/go-prometheus/model"
)

func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	books := []model.Book{}
	rows, err := db.Query("SELECT id, title, author, created_at FROM books")
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt); err != nil {
			return books, err
		}
		books = append(books, b)
	}

	return books, nil
}
