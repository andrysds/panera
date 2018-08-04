package handler

import (
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	message := NewMessage(config.SquadID, update.Message.Text)
	return message
}

func HandleMasterCommand(command string) *tgbotapi.MessageConfig {
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
	message := NewMessage(config.MasterID, result)
	return message
}
