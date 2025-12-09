package handlers

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is the root; if not, show 404
	if r.URL.Path != "/" {
		ServeError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ServeError(w, http.StatusBadRequest)
		return
	}

	// Serve the home page with an empty Result initially
	w.WriteHeader(http.StatusOK)
	renderTemplate(w, indexTmpl, "index", indexData)
}
