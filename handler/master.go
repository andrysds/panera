package handler

import (
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleMasterCommand(chatId int64, command string) *tgbotapi.MessageConfig {
	result := "command is not defined"
	switch command {
	// migrate
	case "standup_init":
		result = migrate.StandupInit()
	case "standup_list_init":
		result = migrate.StandupListInit()

	// standup
	case "standup_new_day":
		result = standup.NewDay()
	}
	message := NewMessage(chatId, result)
	return message
}
