package handler

import (
	"math/rand"
	"strings"

	"github.com/andrysds/panera/helper"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defaultMessage := pickDefaultMessage()
	message := tgbotapi.NewMessage(update.Message.Chat.ID, defaultMessage)

	helper.SendMessage(bot, "Hello World! Panera here.")
}
