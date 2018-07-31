package handler

import (
	"github.com/andrysds/panera/helper"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defaultMessage := "Hello World! Panera here."
	message := tgbotapi.NewMessage(update.Message.Chat.ID, defaultMessage)
	helper.SendMessage(bot, message)
}
