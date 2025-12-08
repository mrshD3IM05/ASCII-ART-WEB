package handlers

import (
	"log"
	"net/http"
)

// serveError renders the `errorTemplate` with dynamic content.
func serveError(w http.ResponseWriter, status int) {
	// Prepare data for the template
	var data ErrorPageData
	switch status {
	case http.StatusNotFound:
		data = ErrNotFound
	case http.StatusBadRequest:
		data = ErrBadRequest
	default:
		data = ErrInternalServer
	}
	// Render the error template

	w.WriteHeader(status)
	if err := errorTmpl.ExecuteTemplate(w, "error", data); err != nil {
		log.Printf("error template execution error: %v", err)
		// Fallback textual response
		http.Error(w, http.StatusText(status), status)
		return
	}
}
