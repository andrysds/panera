package handler

import (
	"gopkg.in/telegram-bot-api.v4"
)

func NewMessage(chatId int64, text string) *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(chatId, text)
	message.ParseMode = "markdown"
	return &message
}
