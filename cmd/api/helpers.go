package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]any

func getIdFromURL(r *http.Request) (int64, error) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idString, 10, 0)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, data envelope, status int, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // limiting body size to 1MB
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(dst)

	var (
		syntaxError           *json.SyntaxError
		unmarshalTypeError    *json.UnmarshalTypeError
		invalidUnmarshalError *json.InvalidUnmarshalError
		maxBytesError         *http.MaxBytesError
	)

	if err != nil {
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character: %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("body contains badly-formed JSON")

		case errors.Is(err, io.EOF):
			return fmt.Errorf("body must not be empty")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field == "" {
				return fmt.Errorf("%s is required", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type for field: %s", unmarshalTypeError.Field)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes(1 MB)", maxBytesError.Limit)
		}
	}

	if err = d.Decode(&struct{}{}); err == nil {
		return fmt.Errorf("body must only contain a single JSON value")
	}

	return nil
}
