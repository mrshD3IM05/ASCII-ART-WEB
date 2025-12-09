package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)

// renderTemplate executes a template and handles errors
func renderTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data pageData) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
		if os.IsNotExist(err) {
			log.Printf("template execution error: %v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("template execution error: %v", err)
		ServeError(w, http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
