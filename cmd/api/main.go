package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string // like: production, development, testing
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	cfg    config       // like PORT, DB, ENV variables, etc
	logger *slog.Logger // app specific logger
}

func main() {
	cfg := config{}

	// server configurations
	flag.IntVar(&cfg.port, "port", 9090, "Port on which the server listens")
	flag.StringVar(&cfg.env, "env", "DEV", "Environment (DEV|PROD|TEST)")

	// database DSN configurations
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("BOOKMYSHOW_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 10*time.Minute, "PostgreSQL max connection idle time")

	flag.Parse()

	// creating a new text logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

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
	if err = s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
