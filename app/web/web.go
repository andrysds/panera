package web

import (
	"log"
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/gorilla/mux"
)

type Web struct {
	*clarity.BasicAuthorizer
}

func NewWeb() *Web {
	return &Web{
		BasicAuthorizer: clarity.NewBasicAuthorizer(
			config.Username,
			config.Password,
		),
	}
}

func (w *Web) Run(started chan<- bool) {
	log.Println("* [web] Listening on tcp://0.0.0.0:" + config.Port)
	started <- true

	router := mux.NewRouter()
	router.HandleFunc("/{command}", w.Handle)
	http.ListenAndServe(":"+config.Port, router)
}
