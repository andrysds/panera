package app

import (
	"log"
	"net/http"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	API     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot() *Bot {
	bot := Bot{API: NewAPI()}
	bot.Updates = NewUpdates(bot.API)
	return &bot
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

func (b *Bot) Run() {
	http.HandleFunc("/", handler.HandleHealthz)
	go http.ListenAndServe(":"+config.Port, nil)

	for update := range b.Updates {
		b.Handle(&update)
	}
}

func (b *Bot) Handle(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	var message *tgbotapi.MessageConfig
	chatID := update.Message.Chat.ID
	handler.LogMessage(update.Message.From.UserName, update.Message.Text)

	switch {
	case b.IsAddedToGroup(update):
		message = handler.HandleGroupInvitation(chatID)
	case update.Message.IsCommand():
		message = handler.HandleCommand(chatID, update.Message.Command())
		message.ReplyToMessageID = update.Message.MessageID
	case chatID == config.MasterID:
		message = handler.HandleMasterMessage(update)
	}
	b.SendMessage(message)
}

func (b *Bot) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		_, err := b.API.Send(message)
		clarity.PrintIfError(err, "error on sending message")
	}
}

func (b *Bot) IsAddedToGroup(update *tgbotapi.Update) bool {
	if update.Message.NewChatMembers != nil {
		members := update.Message.NewChatMembers
		for _, member := range *members {
			if member.UserName == "panera_bot" {
				return true
			}
		}
	}
	return false
}
