package bot

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrysds/clarity"
	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	BotAPI   *tgbotapi.BotAPI
	Updates  tgbotapi.UpdatesChannel
	MasterId int64
	ChatId   int64
}

func NewBot() *Bot {
	botToken := os.Getenv("BOT_TOKEN")
	botAPI := NewBotAPI(botToken)
	updates := NewUpdates(botAPI)
	masterId, _ := strconv.Atoi(os.Getenv("MASTER_ID"))
	chatId, _ := strconv.Atoi(os.Getenv("CHAT_ID"))

	return &Bot{
		BotAPI:   botAPI,
		Updates:  updates,
		MasterId: int64(masterId),
		ChatId:   int64(chatId),
	}
}

func NewBotAPI(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	clarity.PanicIfError(err, "error on creating bot api")
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	return bot
}

func NewUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	webhook := tgbotapi.NewWebhook(webhookUrl + bot.Token)
	_, err := bot.SetWebhook(webhook)
	clarity.PanicIfError(err, "error on setting bot webhook")
	return bot.ListenForWebhook("/" + bot.Token)
}
