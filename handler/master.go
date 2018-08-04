package handler

import (
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleMasterCommand(command string) *tgbotapi.MessageConfig {
	result := NotFoundMessage
	switch command {
	// migrate
	case "init":
		result = migrate.Init()
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

func HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	message := NewMessage(config.SquadID, update.Message.Text)
	return message
}
