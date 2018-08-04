package web

import (
	"net/http"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Route() {
	http.HandleFunc("/healthz", w.HandleHealthz)
}
