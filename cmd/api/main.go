package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string // like: production, development, testing
}

type application struct {
	cfg    config       // like PORT, DB, ENV variables, etc
	logger *slog.Logger // app specific logger
}

func main() {
	cfg := config{}
	flag.IntVar(&cfg.port, "port", 9090, "Port on which the server listens")
	flag.StringVar(&cfg.env, "env", "DEV", "Environment (DEV|PROD|TEST)")
	flag.Parse()

	// creating a new text logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		cfg:    cfg,
		logger: logger,
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  5 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError), // making Server use our logger for Errors
	}

	logger.Info("starting server", "addr", s.Addr, "env", cfg.env)
	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
