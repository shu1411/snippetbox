package main

import (
	"database/sql"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("couldn't load environment variables: %v", err)
	}

	// command flags
	addr := flag.String("addr", ":8080", "HTTP network address")
	dbURL := flag.String("db_url", os.Getenv("DB_URL"), "PostgreSQL database URL")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dbURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)
	logger.Error(http.ListenAndServe(*addr, app.routes()).Error())
	os.Exit(1)
}

func openDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
