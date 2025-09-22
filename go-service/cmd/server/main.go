package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/you/go-service/internal/metrics"
)

func main() {
	r := chi.NewRouter()

	// Add useful middleware
	r.Use(middleware.Logger)    // logs each request
	r.Use(middleware.Recoverer) // recovers from panics

	// Routes
	r.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		m := metrics.Collect()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	})

	// Start server
	http.ListenAndServe(":8080", r)
}
