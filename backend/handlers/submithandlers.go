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
		// Success: render the template with ASCII art and echo the inputs
		escapedASCII := html.EscapeString(asciiArt)
		data := struct {
			Title    string
			Result   template.HTML
			RawInput string
			Font     string
		}{
			Title:    "ASCII Art Web",
			Result:   template.HTML("<pre>" + escapedASCII + "</pre>"),
			RawInput: userText,
			Font:     fontName,
		}
		w.WriteHeader(http.StatusOK)
		renderTemplate(w, resultTmpl, "result", data)

	case http.StatusNotFound, http.StatusBadRequest:
		// Render an error page for known client errors
		serveError(w, statusCode)

	default:
		// Unexpected status code -> show generic server error
		serveError(w, http.StatusInternalServerError)
	}
}
