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
	PageTitle        string
	FormAction       string
	PastAction       string
	PastActionStatus string
}

func (t *templateData) setPastActionInfo(r *http.Request) {
	t.PastAction = r.URL.Query().Get("past_action")
	t.PastActionStatus = r.URL.Query().Get("past_action_status")
}

type afterActionOpts struct {
	action       string
	err          error
	redirectPath string
}

func afterAction(w http.ResponseWriter, r *http.Request, opts afterActionOpts) {
	if opts.err != nil {
		internalServerError(w, opts.err)
	} else {
		status := "success"
		if opts.err != nil {
			status = "danger"
		}
		redirectBack(w, r, opts.redirectPath+"?past_action="+opts.action+"&past_action_status="+status)
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
	http.Redirect(w, r, path, 302)
}
