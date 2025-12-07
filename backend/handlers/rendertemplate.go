package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// renderTemplate executes a template and handles errors
func renderTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data interface{}) {
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("template execution error: %v", err)
		serveError(w, http.StatusInternalServerError)
	}
}
