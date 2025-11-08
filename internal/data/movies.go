package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Runtime   int32     `json:"runtime"`
	Year      int32     `json:"year"`
	Genres    []string  `json:"genres"`
	Storyline string    `json:"storyline"`
	Version   int32     `json:"version"`
	Languages []string  `json:"languages"`
}
