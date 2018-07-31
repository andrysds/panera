package app

import (
	"log"
	"os"

	"github.com/andrysds/pnr_bot/handler"
	"github.com/andrysds/pnr_bot/helper"
	"gopkg.in/telegram-bot-api.v4"
)

type PnrBot struct {
	Bot     *tgbotapi.BotAPI
	Updates <-chan tgbotapi.Update
}

func NewPnrBot() *PnrBot {
	botToken := os.Getenv("BOT_TOKEN")
	bot := NewBot(botToken)

	updates := NewUpdates(bot)

	return &PnrBot{
		Bot:     bot,
		Updates: updates,
	}
}

func NewBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	helper.CheckAndFatal(err)
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func NewUpdates(bot *tgbotapi.BotAPI) <-chan tgbotapi.Update {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	webhook := tgbotapi.NewWebhook(webhookUrl + bot.Token)
	_, err := bot.SetWebhook(webhook)
	helper.CheckAndFatal(err)
	return bot.ListenForWebhook("/" + bot.Token)
}

func (b *PnrBot) Run() {
	for update := range b.Updates {
		if update.Message == nil {
			continue
		}
		helper.LogMessage(update.Message)

		switch update.Message.Text {
		default:
			handler.HandleDefault(b.Bot, update)
		}
	}
}
