package bot

import (
	"log"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	API     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
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

func (b *Bot) Run(started chan<- bool) {
	log.Println("* [bot] Listening from webhook")
	started <- true

	for update := range b.Updates {
		b.Handle(&update)
	}
}
