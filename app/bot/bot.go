package bot

import (
	"log"
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"github.com/newrelic/go-agent"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	API      *tgbotapi.BotAPI
	Updates  tgbotapi.UpdatesChannel
	NewRelic newrelic.Application
}

func NewBot() *Bot {
	API := NewAPI()
	updates := NewUpdates(API)
	return &Bot{
		API:     API,
		Updates: updates,
	}
}

func NewAPI() *tgbotapi.BotAPI {
	if config.BotToken != "" {
		API, err := tgbotapi.NewBotAPI(config.BotToken)
		clarity.PanicIfError(err, "error on creating bot api")
		log.Printf("* Authorized on account %s\n", API.Self.UserName)
		return API
	}
	return nil
}

func NewUpdates(API *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	if API != nil {
		webhook := tgbotapi.NewWebhook(config.WebhookUrl + API.Token)
		_, err := API.SetWebhook(webhook)
		clarity.PanicIfError(err, "error on setting bot webhook")
		return API.ListenForWebhook("/" + API.Token)
	}
	return nil
}

func NewNewRelic() newrelic.Application {
	config := newrelic.NewConfig("Panera", config.NewRelicKey)
	if app, err := newrelic.NewApplication(config); err == nil {
		return app
	} else {
		return nil
	}
}

func (b *Bot) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc(newrelic.WrapHandleFunc(b.NewRelic, "/healthz", handler.HandleHealthz))
	go http.ListenAndServe(":"+config.Port, mux)

	for update := range b.Updates {
		b.Handle(&update)
	}
}
