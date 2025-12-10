package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Handler initializes all HTTP routes and parses templates.
// Call this once during server startup
func Handler() {

	// Parse all templates
	mustParse := func(files ...string) *template.Template {
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("error template execution error: %v", err)
				// Fallback textual response
				os.Exit(1)
			}
			log.Fatalf("failed to parse templates %v: %v", files, err)
		}
		return tmpl
	}
	templates = mustParse("frontend/common.html", "frontend/index.html", "frontend/result.html", "frontend/error.html")
	// Register HTTP routes
	http.HandleFunc("/", indexHandler)           // GET requests to / serve the home page
	http.HandleFunc("/ascii-art", submitHandler) // POST requests to /ascii-art generate ASCII art (server-side)
}
