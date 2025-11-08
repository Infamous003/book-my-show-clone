package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// routes returns a Handler that has all the app specific routes registered
func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", app.rootHandler)
	r.Get("/healthcheck", app.healthcheckHandler)
	r.Get("/movies/{id}", app.getMovieHandler)

	return r
}

func (app *application) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("What's up!\n"))
}
