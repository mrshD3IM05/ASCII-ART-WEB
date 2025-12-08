package handlers

import (
	"html/template"
	"net/http"
)

var indexTmpl *template.Template

var errorTmpl *template.Template

var resultTmpl *template.Template

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
