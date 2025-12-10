package handlers

import (
	"bytes"
	"log"
	"net/http"
)

// renderTemplate executes a template and handles errors
// renderTemplate executes a specific template from the global set and handles errors
func renderTemplate(w http.ResponseWriter, name string, data pageData) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("template execution error: %v", err)
		ServeError(w, http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
