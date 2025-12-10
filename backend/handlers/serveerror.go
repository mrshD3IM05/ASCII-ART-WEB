package handlers

import (
	"net/http"
)

// ServeError renders the `errorTemplate` with dynamic content.
func ServeError(w http.ResponseWriter, status int) {
	// Prepare data for the template
	var data pageData
	switch status {
	case http.StatusNotFound:
		data = ErrNotFound
	case http.StatusBadRequest:
		data = ErrBadRequest
	default:
		data = ErrInternalServer
	}
	w.WriteHeader(status)
	renderTemplate(w, "error", data)
}
