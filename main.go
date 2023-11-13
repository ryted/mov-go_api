package main

import (
	// Custom
	MovieRoutes "github.com/ryted/mov-go_api/src/Router"

	// Built-In
	"fmt"
	"log"
	"net/http"

	// Third-Parties
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Router Instance
	r := mux.NewRouter()

	// Movie Routes
	for _, route := range MovieRoutes.MovieRoutes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}

	fmt.Println("Server berjalan di port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
