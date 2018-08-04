package bot

import (
	"log"
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
	masterId, _ := strconv.Atoi(os.Getenv("MASTER_ID"))
	chatId, _ := strconv.Atoi(os.Getenv("CHAT_ID"))

	bot := &Bot{
		MasterId: int64(masterId),
		ChatId:   int64(chatId),
	}

	return bot
}

func NewBotAPI() *tgbotapi.BotAPI {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken != "" {
		botAPI, err := tgbotapi.NewBotAPI(botToken)
		clarity.PanicIfError(err, "error on creating bot api")
		log.Printf("* Authorized on account %s\n", botAPI.Self.UserName)
		return botAPI
	}
	return nil
}

func NewUpdates(botAPI *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	if botAPI != nil {
		webhookUrl := os.Getenv("WEBHOOK_URL")
		webhook := tgbotapi.NewWebhook(webhookUrl + botAPI.Token)
		_, err := botAPI.SetWebhook(webhook)
		clarity.PanicIfError(err, "error on setting bot webhook")
		return botAPI.ListenForWebhook("/" + botAPI.Token)
	}
	return nil
}

func (b *Bot) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		_, err := b.BotAPI.Send(message)
		clarity.PrintIfError(err, "error on sending message")
	}
}

func (b *Bot) LogMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s\n", message.From.UserName, message.Text)
}
