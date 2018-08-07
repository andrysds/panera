package web

import (
	"fmt"
	"net/http"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
)

func (w *Web) Handle(wr http.ResponseWriter, r *http.Request) {
	var message *tgbotapi.MessageConfig
	command := mux.Vars(r)["command"]
	handler.Log(r.Method, "/"+command)

	switch command {
	case "healthz":
		message = handler.NewMessage(config.MasterID, handler.OKMessage)
	default:
		if err := w.Authorize(r.Header); err == nil {
			message = handler.HandleCommand(config.MasterID, command)
		} else {
			message = handler.NewMessage(config.MasterID, handler.UnauthorizedMessage)
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
		w.SendMessage(message)
	}
}

func (w *Web) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		handler.Log("panera", message.Text)
	}
}
