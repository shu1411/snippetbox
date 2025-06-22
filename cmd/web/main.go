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

	// set up server and route declarations
	mux := app.routes()

	// run the server and log any info and errors
	logger.Info("starting server", "addr", *addr)
	logger.Error(http.ListenAndServe(*addr, mux).Error())

	os.Exit(1)
}
