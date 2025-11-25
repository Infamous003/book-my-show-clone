package data

import (
	"time"

	"github.com/Infamous003/book-my-show-clone/internal/data/validator"
)

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

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 100, "title", "must not be more than 500 bytes long")

	v.Check(movie.Description != "", "description", "must be provided")
	v.Check(len(movie.Description) <= 1000, "description", "must not be more than 1000 bytes")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must not be greater than 1888")
	v.Check(movie.Year < int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) > 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicates")

	v.Check(movie.Languages != nil, "languages", "must be provided")
	v.Check(len(movie.Languages) > 1, "languages", "must contain at least 1 language")
	v.Check(len(movie.Languages) <= 10, "languages", "must not contain more than 10 languages")
	v.Check(validator.Unique(movie.Languages), "languages", "must not contain duplicates")

}
