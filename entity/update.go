package entity

import (
	"gopkg.in/telegram-bot-api.v4"
)

var BlankUpdate = &tgbotapi.Update{
	Message: &tgbotapi.Message{
		Chat: &tgbotapi.Chat{},
		From: &tgbotapi.User{},
	},
}
