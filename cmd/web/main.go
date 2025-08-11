package main

import (
	"log"
	"net/http"
)

const port = 8080

func main() {
	mux := http.NewServeMux()

	// restrict this route to exact matches on '/' only
	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on :%d", port)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
