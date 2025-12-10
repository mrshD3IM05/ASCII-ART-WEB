package handlers

import (
	"ascii-art-web/backend/ASCII"
	"html"
	"html/template"
	"net/http"
)

// submitHandler processes POST requests to generate ASCII art.
// It expects form data with "text" and "font" fields.
func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Validate HTTP method
	if r.Method != http.MethodPost {
		ServeError(w, http.StatusBadRequest)
		return
	}

	// Parse form data from the request
	if err := r.ParseForm(); err != nil {
		ServeError(w, http.StatusBadRequest)
		return
	}

	// Extract form values
	userText := r.FormValue("text")
	fontName := r.FormValue("font")

	// Generate ASCII art using the selected font
	asciiArt, statusCode := ASCII.CreateASCIIArtTable(userText, fontName)

	// Handle ASCII art generation response using a switch that maps statuses to error pages
	switch statusCode {
	case http.StatusOK:
		// Success: render the template with ASCII art and echo the inputs
		escapedASCII := html.EscapeString(asciiArt)
		data := pageData{
			Title:  "ASCII Art Web - Result",
			Result: template.HTML("<pre>" + escapedASCII + "</pre>"),
		}
		w.WriteHeader(http.StatusOK)
		renderTemplate(w, "result", data)

	case http.StatusNotFound, http.StatusBadRequest, http.StatusInternalServerError:
		// Render an error page for known client errors
		ServeError(w, statusCode)

	default:
		// Unexpected status code -> show generic server error
		ServeError(w, http.StatusInternalServerError)
	}
}
