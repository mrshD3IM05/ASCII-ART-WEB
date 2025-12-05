package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Handler initializes all HTTP routes and parses templates.
// Call this once during server startup
func Handler() {
	var err error

	// Parse the main index page template
	indexTemplate, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("failed to parse index template: %v", err)
	}

	// Parse the generic error page template
	errorTemplate, err = template.ParseFiles("templates/error.html")
	if err != nil {
		log.Fatalf("failed to parse error template: %v", err)
	}

	// Register HTTP routes
	http.HandleFunc("/", indexHandler)           // GET requests to / serve the home page
	http.HandleFunc("/ascii-art", submitHandler) // POST requests to /ascii-art generate ASCII art
}
