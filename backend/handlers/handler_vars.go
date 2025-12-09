package handlers

import (
	"html/template"
	"net/http"
)

// indexTmpl holds the parsed templates for the index page
var indexTmpl *template.Template

// errorTmpl holds the parsed templates for the error page
var errorTmpl *template.Template

// resultTmpl holds the parsed templates for the result page
var resultTmpl *template.Template

type pageData struct {
	Status  int
	Title   string
	Message string
	Detail  string
	Result  template.HTML
}

var (
	indexData = pageData{
		Title: "ASCII Art Web",
	}
	ErrNotFound = pageData{
		Status:  http.StatusNotFound,
		Title:   "Not Found",
		Message: "The requested resource was not found.",
	}

	ErrInternalServer = pageData{
		Status:  http.StatusInternalServerError,
		Title:   "Internal Server Error",
		Message: "An unexpected error occurred on the server.",
	}

	ErrBadRequest = pageData{
		Status:  http.StatusBadRequest,
		Title:   "Bad Request",
		Message: "The server could not understand the request due to invalid syntax.",
	}
)
