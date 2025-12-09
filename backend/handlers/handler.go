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

	// Helper function to parse templates
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

	// Parse templates
	errorTmpl = mustParse("frontend/App.html", "frontend/pages/error.html")
	indexTmpl = mustParse("frontend/App.html", "frontend/pages/index.html")
	resultTmpl = mustParse("frontend/App.html", "frontend/pages/result.html")

	// Register HTTP routes
	http.HandleFunc("/", indexHandler)           // GET requests to / serve the home page
	http.HandleFunc("/ascii-art", submitHandler) // POST requests to /ascii-art generate ASCII art (server-side)
}
