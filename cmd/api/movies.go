package main

import (
	"net/http"
	"time"

	"github.com/Infamous003/book-my-show-clone/internal/data"
)

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromURL(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Puss in Boots",
		CreatedAt: time.Now(),
		Year:      2011,
		Runtime:   107,
		Genres:    []string{"Animation", "Adventure", "Fantasy"},
		Storyline: "A story about a criminal cat",
		Version:   1,
		Languages: []string{"English", "French", "Spanish", "Hindi"},
	}

	headers := http.Header{}
	headers.Set("Languages", "en")

	if err = app.writeJSON(w, movie, 200, headers); err != nil {
		w.Write([]byte("the server encountered an error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
