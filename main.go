package main

import (
	"log"
	"net/http"
)

const port = 8080

func main() {
	mux := http.NewServeMux()
	
	// restrict this route to exact matches on '/' only
	mux.HandleFunc("/{$}", home)

	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on :%d", port)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

