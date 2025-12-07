package handlers

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is the root; if not, show 404
	if r.URL.Path != "/" {
		serveError(w, http.StatusNotFound)
		return
	}

	// Serve the home page with an empty Result initially
	data := struct {
		Title    string
		Result   template.HTML
		RawInput string
		Font     string
	}{
		Title: "ASCII Art Web",
	}
	w.WriteHeader(http.StatusOK)
	renderTemplate(w, indexTmpl, "index", data)
}
