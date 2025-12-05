package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is the root; if not, show 404
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		serveError(w, http.StatusNotFound)
		return
	}

	// Serve the home page with an empty Result initially
	data := struct{ Result template.HTML }{}
	w.WriteHeader(http.StatusOK)
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Printf("template execution error: %v", err)
		serveError(w, http.StatusInternalServerError)
		return
	}
}
