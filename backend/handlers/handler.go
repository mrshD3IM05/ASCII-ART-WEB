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
	indexTmpl, err = template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		log.Fatalf("failed to parse index templates: %v", err)
	}

	// Parse the error page template
	errorTmpl, err = template.ParseFiles("templates/base.html", "templates/error.html")
	if err != nil {
		log.Fatalf("failed to parse error templates: %v", err)
	}

	// Parse the result page template
	resultTmpl, err = template.ParseFiles("templates/base.html", "templates/result.html")
	if err != nil {
		log.Fatalf("failed to parse result templates: %v", err)
	}

	// Register HTTP routes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.HandleFunc("/", indexHandler)           // GET requests to / serve the home page
	http.HandleFunc("/ascii-art", submitHandler) // POST requests to /ascii-art generate ASCII art (server-side)

}
