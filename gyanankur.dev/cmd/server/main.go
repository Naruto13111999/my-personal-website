package main

import (
	"log"
	"net/http"
	"os"

	"gyanankur.dev/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h, err := handlers.New()
	if err != nil {
		log.Fatalf("failed to initialize handlers: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.Index)
	mux.HandleFunc("GET /api/profile", h.APIProfile)
	mux.HandleFunc("GET /health", h.Health)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Printf("gyanankur.dev running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
