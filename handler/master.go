package handler

import (
	"strconv"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleMasterCommand(command string) *tgbotapi.MessageConfig {
	result := entity.NotFoundMessage
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
		result = entity.NewDayStandup()
		if result == entity.NotFoundMessage {
			result = entity.NewPeriodStandupList()
		}
	case "standup_new_period":
		result = entity.NewPeriodStandupList()
	}
	message := entity.NewMessage(config.MasterID, result)
	return message
}

func HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	var message *tgbotapi.MessageConfig
	if update.Message.ForwardFrom != nil {
		messageText := update.Message.ForwardFrom.UserName + " "
		messageText += strconv.Itoa(update.Message.ForwardFrom.ID)
		message = entity.NewMessage(config.MasterID, messageText)
	} else {
		message = entity.NewMessage(config.SquadID, update.Message.Text)
	}
	return message
}
