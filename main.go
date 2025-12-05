package main

import (
	"ascii-art-web/backend/handlers"
	"log"
	"net/http"
)

// main is the entry point for the ASCII Art Web server.
// It initializes routes and starts listening on port 8080.
func main() {
	// Initialize HTTP routes and parse templates
	handlers.Handler()

	// Start the HTTP server on port 8080
	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
