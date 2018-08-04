package handler

import (
	"log"

	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

func Log(caller, message string) {
	log.Printf("[%s] %s\n", caller, message)
}

func HandleCommand(chatID int64, command string) *tgbotapi.MessageConfig {
	var message *tgbotapi.MessageConfig
	switch command {
	case "standup":
		message = HandleStandup(chatID)
	case "standup_list":
		message = HandleStandupList(chatID)
	case "standup_skip":
		message = HandleStandupSkip(chatID)
	default:
		if chatID == config.MasterID {
			message = HandleMasterCommand(command)
		}
	}
	return message
}
