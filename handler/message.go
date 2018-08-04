package handler

import (
	"gopkg.in/telegram-bot-api.v4"
)

func NewMessage(chatID int64, text string) *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = "markdown"
	return &message
}
