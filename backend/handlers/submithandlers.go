package handlers

import (
	"ascii-art-web/backend/ASCII"
	"html"
	"html/template"
	"log"
	"net/http"
)

// submitHandler processes POST requests to generate ASCII art.
// It expects form data with "text" and "font" fields.
func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Validate HTTP method
	if r.Method != http.MethodPost {
		serveError(w, http.StatusBadRequest)
		return
	}

	// Parse form data from the request
	if err := r.ParseForm(); err != nil {
		serveError(w, http.StatusBadRequest)
		return
	}

	// Extract form values
	userText := r.FormValue("text")
	fontName := r.FormValue("font")

	// Generate ASCII art using the selected font
	asciiArt, statusCode, _ := ASCII.CreateASCIIArtTable(userText, fontName)

	// Handle ASCII art generation response using a switch that maps statuses to error pages
	switch statusCode {
	case http.StatusOK:
		// Success: render the template with ASCII art
		escapedASCII := html.EscapeString(asciiArt)
		data := struct{ Result template.HTML }{
			Result: template.HTML("<pre>" + escapedASCII + "</pre>"),
		}
		w.WriteHeader(http.StatusOK)
		if err := indexTemplate.Execute(w, data); err != nil {
			log.Printf("template execution error: %v", err)
			serveError(w, http.StatusInternalServerError)
			return
		}

	case http.StatusNotFound, http.StatusBadRequest:
		// Font not found
		serveError(w, statusCode)
	default:
		// Unexpected status code
		serveError(w, http.StatusInternalServerError)
	}
}
