package handler

import (
	"io"
	"net/http"

	"github.com/andrysds/panera/template"
)

// Healthz is handler function for GET /healthz
func Healthz(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "OK!\n")
}

// NotFound is handler function for 404
func NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderError(w, "404 page not found")
}

// Index is handler function for GET /
func Index(w http.ResponseWriter, _ *http.Request) {
	template.Execute(w, "index.html", nil)
}

type templateData struct {
	PageTitle  string
	FormAction string
}

func afterAction(w http.ResponseWriter, r *http.Request, err error, redirectPath string) {
	if err != nil {
		internalServerError(w, err)
	} else {
		redirectBack(w, r, redirectPath)
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	renderError(w, err.Error())
}

func renderError(w http.ResponseWriter, msg string) {
	template.Template.ExecuteTemplate(w, "error.html", msg)
}

func redirectBack(w http.ResponseWriter, r *http.Request, path string) {
	http.Redirect(w, r, path, 301)
}
