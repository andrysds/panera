package web

import (
	"log"
	"net/http"
	"os"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Run(started chan<- bool) {
	w.Route()
	port := ":" + os.Getenv("PORT")
	log.Println("* [web] Listening on tcp://0.0.0.0" + port)
	started <- true
	http.ListenAndServe(port, nil)
}

func (w *Web) Route() {
	http.HandleFunc("/healthz", w.HandleHealthz)
}
