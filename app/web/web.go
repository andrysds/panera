package web

import (
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
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
	handler.Handler.HandleFunc("/{command}", w.Handle)
	http.ListenAndServe(":"+config.Port, handler.Handler)
}
