package entity

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	OKMessage       = "ok"        // 200
	NotFoundMessage = "not found" // 404
)

func NewMessage(update *tgbotapi.Update, text string) *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	message.ParseMode = "markdown"
	message.ReplyToMessageID = update.Message.MessageID
	return &message
}

func LogMessage(caller, message string) {
	log.Printf("[%s] %s\n", caller, message)
}
