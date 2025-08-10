package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display a form for creating a new snippet...")
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Save a new snippet...")
}
