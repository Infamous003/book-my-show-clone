package data

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Runtime     Runtime   `json:"runtime"`
	Year        int32     `json:"year"`
	Genres      []string  `json:"genres"`
	Languages   []string  `json:"languages"`
	Version     int32     `json:"version"`
	UpdatedAt   time.Time `json:"-"`
}
