package app

import (
	"log"
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	BotAPI  entity.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot() *Bot {
	botAPI, err := tgbotapi.NewBotAPI(config.BotToken)
	clarity.PanicIfError(err, "error on creating bot api")
	log.Printf("* Authorized on account %s\n", botAPI.Self.UserName)

	webhook := tgbotapi.NewWebhook(config.WebhookUrl + botAPI.Token)
	_, err = botAPI.SetWebhook(webhook)
	clarity.PanicIfError(err, "error on setting bot webhook")
	updates := botAPI.ListenForWebhook("/" + botAPI.Token)

	bot := Bot{
		BotAPI:  botAPI,
		Updates: updates,
	}
	return &bot
}

func (b *Bot) GetBotAPI() entity.BotAPI {
	return b.BotAPI
}

func (b *Bot) Run() {
	http.HandleFunc("/", handler.HandleHealthz)
	go http.ListenAndServe(":"+config.Port, nil)
	log.Println("* Listening on tcp://0.0.0.0:" + config.Port)
	log.Println("Use Ctrl-C to stop")

	for update := range b.Updates {
		handler.HandleUpdate(&update, b.BotAPI)
	}
}
