package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, status int, message string) {
	js := envelope{"error": message}

	err := app.writeJSON(w, js, status, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter) {
	message := "the requested resource could not be found"
	app.errorResponse(w, http.StatusNotFound, message)
}

func (app *application) serverErrorResponse(w http.ResponseWriter) {
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, http.StatusInternalServerError, message)
}
