package handler

import (
	"log"
	"strconv"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
)

var Handler *mux.Router

func Init() {
	Handler = mux.NewRouter()
	Handler.HandleFunc("/healthz", HandleHealthz)

	log.Println("* Handler initialized")
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

func HandleGroupInvitation(chatID int64) *tgbotapi.MessageConfig {
	messageText := "I was invited to " + strconv.FormatInt(chatID, 10)
	message := entity.NewMessage(config.MasterID, messageText)
	return message
}

func LogMessage(caller, message string) {
	log.Printf("[%s] %s\n", caller, message)
}
