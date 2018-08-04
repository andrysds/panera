package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrysds/clarity"
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
	clarity.PanicIfError(err, "error on creating bot api")
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	return bot
}

func NewUpdates(bot *tgbotapi.BotAPI) <-chan tgbotapi.Update {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	webhook := tgbotapi.NewWebhook(webhookUrl + bot.Token)
	_, err := bot.SetWebhook(webhook)
	clarity.PanicIfError(err, "error on setting bot webhook")
	return bot.ListenForWebhook("/" + bot.Token)
}
