package handler

import (
	"log"
	"strconv"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleCommand(chatID int64, command string, bot *tgbotapi.BotAPI) *tgbotapi.MessageConfig {
	var message *tgbotapi.MessageConfig
	switch command {
	case "standup":
		message = HandleStandup(chatID)
	case "standup_list":
		message = HandleStandupList(chatID)
	case "standup_skip":
		message = HandleStandupSkip(chatID)
	case "birthdays":
		message = HandleBirthdays(chatID)
	default:
		if chatID == config.MasterID {
			message = HandleMasterCommand(command, bot)
		}
	}
	return message
}

func HandleGroupInvitation(chatID int64) *tgbotapi.MessageConfig {
	messageText := "I was invited to " + strconv.FormatInt(chatID, 10)
	message := entity.NewMessage(config.MasterID, messageText)
	return message
}

func LogMessage(caller, message string) {
	log.Printf("[%s] %s\n", caller, message)
}
