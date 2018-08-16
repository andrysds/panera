package web

import (
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
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

func (w *Web) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.HandleHealthz)
	router.HandleFunc("/{command}", w.Handle)
	http.ListenAndServe(":"+config.Port, router)
}
