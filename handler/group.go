package handler

import (
	"strconv"

	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleGroupInvitation(chatID int64) *tgbotapi.MessageConfig {
	messageText := "I was invited to " + strconv.FormatInt(chatID, 10)
	message := NewMessage(config.MasterID, messageText)
	return message
}
