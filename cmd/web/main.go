package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// set up command-line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// set up logger, using nil for default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// set up app for handler dependencies
	app := &application{
		logger: logger,
	}

	// set up Server
	mux := http.NewServeMux()

	// file server and handlers
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// logging
	logger.Info("starting server", "addr", *addr)
	logger.Error(http.ListenAndServe(*addr, mux).Error())

	os.Exit(1)
}
