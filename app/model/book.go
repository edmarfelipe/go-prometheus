package model

import "time"

type Book struct {
	ID        int
	Title     string
	Author    string
	CreatedAt time.Time
}
