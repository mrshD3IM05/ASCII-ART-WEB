package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Handler initializes all HTTP routes and parses templates.
// Call this once during server startup
func Handler() {

	// Helper function to parse templates
	mustParse := func(files ...string) *template.Template {
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Fatalf("failed to parse templates %v: %v", files, err)
		}
		return tmpl
	}

	// Parse templates
	indexTmpl = mustParse("templates/App.html", "templates/pages/index.html")
	errorTmpl = mustParse("templates/App.html", "templates/pages/error.html")
	resultTmpl = mustParse("templates/App.html", "templates/pages/result.html")

	// Register HTTP routes
	http.HandleFunc("/", indexHandler)           // GET requests to / serve the home page
	http.HandleFunc("/ascii-art", submitHandler) // POST requests to /ascii-art generate ASCII art (server-side)
}
