package main

import (
	"net/http"
	"time"

	"github.com/Infamous003/book-my-show-clone/internal/data"
)

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromURL(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:          id,
		Title:       "Puss in Boots",
		CreatedAt:   time.Now(),
		Description: "A story about a criminal cat",
		Year:        2011,
		Runtime:     107,
		Genres:      []string{"Animation", "Adventure", "Fantasy"},
		Version:     1,
		Languages:   []string{"English", "French", "Spanish", "Hindi"},
		UpdatedAt:   time.Now(),
	}

	headers := http.Header{}
	headers.Set("Languages", "en")

	if err = app.writeJSON(w, envelope{"movie": movie}, 200, headers); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// creating a placeholder struct to store incoming expected values
	var input struct {
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Year        int32        `json:"year"`
		Runtime     data.Runtime `json:"runtime"`
		Genres      []*string    `json:"genres"`
		Languages   []*string    `json:"languages"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, envelope{"movie": input}, http.StatusCreated, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
