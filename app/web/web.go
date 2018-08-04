package web

import (
	"log"
	"net/http"

	"github.com/andrysds/panera/config"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Run(started chan<- bool) {
	w.Route()
	log.Println("* [web] Listening on tcp://0.0.0.0" + config.Port)
	started <- true
	http.ListenAndServe(config.Port, nil)
}

func (w *Web) Route() {
	http.HandleFunc("/healthz", w.Healthz)
}
