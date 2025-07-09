package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shu1411/snippetbox/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// set up command-line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:webtemp@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// set up logger, using nil for default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// open database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// set up app for handler dependencies
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	// set up server and route declarations
	mux := app.routes()

	// run the server and log any info and errors
	logger.Info("starting server", "addr", *addr)
	logger.Error(http.ListenAndServe(*addr, mux).Error())

	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
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
