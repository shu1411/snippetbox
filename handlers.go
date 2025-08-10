package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Snippetbox")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display a specific snippet...")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display a form for creating a new snippet...")
}
