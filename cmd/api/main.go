package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("What's up!"))
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  2 * time.Hour,
	}
	fmt.Println("Starting server on PORT 8080")

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
