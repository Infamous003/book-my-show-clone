package main

import (
	"fmt"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	js := envelope{"error": message}

	err := app.writeJSON(w, js, status, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	app.logger.Error(err.Error())
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s is not allowed for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
