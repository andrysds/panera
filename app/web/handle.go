package web

import (
	"fmt"
	"net/http"

	"github.com/andrysds/panera/handler"
	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
)

func (w *Web) Handle(wr http.ResponseWriter, r *http.Request) {
	var message *tgbotapi.MessageConfig
	command := mux.Vars(r)["command"]

	switch command {
	case "healthz":
		message = handler.NewMessage(0, handler.OKMessage)
	default:
		if err := w.BasicAuthorizer.Authorize(r.Header); err == nil {
			message = handler.HandleCommand(0, command)
		} else {
			message = handler.NewMessage(0, handler.UnauthorizedMessage)

		}
	}

	if message != nil {
		switch message.Text {
		case handler.NotFoundMessage:
			wr.WriteHeader(http.StatusNotFound)
		case handler.UnauthorizedMessage:
			wr.WriteHeader(http.StatusUnauthorized)
		}
		fmt.Fprintf(wr, message.Text)
	}
}
