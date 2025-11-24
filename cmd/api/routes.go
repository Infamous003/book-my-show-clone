package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes returns a Handler that has all the app specific routes registered
func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	// plugging in middlewares
	r.Use(app.recoverPanic)
	r.Use(middleware.Logger)

	// custom errors for 404 and 405
	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/api/v1/healthcheck", app.healthcheckHandler)
	r.Get("/api/v1/movies/{id}", app.getMovieHandler)
	r.Post("/api/v1/movies", app.createMovieHandler)

	return r
}
