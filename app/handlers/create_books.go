package handlers

import (
	"database/sql"
	"net/http"

	"github.com/edmarfelipe/go-prometheus/model"
)

func CreateBooks(db *sql.DB) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		db.Query("")

		result := model.Book{}

		respondWithJSON(w, http.StatusOK, result)
	}
}
