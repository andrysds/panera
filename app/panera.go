package app

import (
	"log"
	"os"
	"strconv"

	"github.com/andrysds/panera/handler"
	"github.com/andrysds/panera/helper"
	"gopkg.in/telegram-bot-api.v4"
)

type Panera struct {
	Bot      *tgbotapi.BotAPI
	Updates  <-chan tgbotapi.Update
	MasterId int64
	ChatId   int64
}

func NewPanera() *Panera {
	botToken := os.Getenv("BOT_TOKEN")
	bot := NewBot(botToken)
	updates := NewUpdates(bot)
	masterId, _ := strconv.Atoi(os.Getenv("MASTER_ID"))
	chatId, _ := strconv.Atoi(os.Getenv("CHAT_ID"))

	return &Panera{
		Bot:      bot,
		Updates:  updates,
		MasterId: int64(masterId),
		ChatId:   int64(chatId),
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

func (b *Panera) Run() {
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
