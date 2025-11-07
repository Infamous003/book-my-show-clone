package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: available\n"))
	w.Write([]byte("environment: development\n"))
	w.Write([]byte("version: 1\n"))
}
