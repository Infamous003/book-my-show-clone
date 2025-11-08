package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"status":      "available",
		"environment": "development",
		"version":     1,
	}

	app.writeJSON(w, data, 200, nil)
}
