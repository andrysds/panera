package handler

import (
	"strconv"
	"strings"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleMasterCommand(command string, bot *tgbotapi.BotAPI) *tgbotapi.MessageConfig {
	result := entity.NotFoundMessage
	switch command {
	// birthday
	case "birthday_kick":
		result = HandleBirthdayKick(bot)
	case "birthday_unban":
		result = HandleBirthdayUnban(bot)

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
		messageText := strconv.Itoa(update.Message.ForwardFrom.ID)
		message = entity.NewMessage(config.MasterID, messageText)
	} else {
		messageText := strings.Replace(update.Message.Text, "<bq>", "`", -1)
		message = entity.NewMessage(config.SquadID, messageText)
	}
	return message
}
