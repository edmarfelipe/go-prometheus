package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edmarfelipe/go-prometheus/database"
)

func GetBooks(db *sql.DB) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := database.GetAllBooks(db)
		if err != nil {
			log.Printf("Error: %s", err)
			respondWithError(w, 500, "Internal Error")
			return
		}

		respondWithJSON(w, http.StatusOK, books)
	}
}
