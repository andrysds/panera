package entity

import (
	"gopkg.in/telegram-bot-api.v4"
)

const (
	OKMessage       = "ok"        // 200
	NotFoundMessage = "not found" // 404
)

func NewMessage(chatID int64, text string) *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = "markdown"
	return &message
}
