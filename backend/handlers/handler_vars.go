package handlers

import (
	"html/template"
	"net/http"
)

// indexTemplate is the main template for the home page
var indexTemplate *template.Template

// errorTemplate is the template for all non-200 error pages
var errorTemplate *template.Template

type ErrorPageData struct {
	Status  int
	Title   string
	Message string
	Detail  string
}

var (
	ErrNotFound = ErrorPageData{
		Status:  http.StatusNotFound,
		Title:   "Not Found",
		Message: "The requested resource was not found.",
	}

	ErrInternalServer = ErrorPageData{
		Status:  http.StatusInternalServerError,
		Title:   "Internal Server Error",
		Message: "An unexpected error occurred on the server.",
	}

	ErrBadRequest = ErrorPageData{
		Status:  http.StatusBadRequest,
		Title:   "Bad Request",
		Message: "The server could not understand the request due to invalid syntax.",
	}
)
