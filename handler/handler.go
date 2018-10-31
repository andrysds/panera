package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleUpdate(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if update.Message == nil {
		return
	}
	entity.LogMessage(update.Message.From.UserName, update.Message.Text)

	switch {
	case IsAddedToGroup(update):
		HandleGroupInvitation(update, botAPI)

	case update.Message.IsCommand():
		switch update.Message.Command() {
		// birthday
		case "birthdays":
			HandleBirthdays(update, botAPI)
		case "birthday_kick":
			HandleBirthdayKick(update, botAPI)
		case "birthday_link":
			HandleBirthdayLink(update, botAPI)

		// standup
		case "standup":
			HandleStandup(update, botAPI)
		case "standup_list":
			HandleStandupList(update, botAPI)
		case "standup_new_day":
			HandleStandupNewDay(update, botAPI)
		case "standup_new_period":
			HandleStandupNewPeriod(update, botAPI)
		case "standup_skip":
			HandleStandupSkip(update, botAPI)
		}

	default:
		HandleMasterMessage(update, botAPI)
	}
}

func HandleMasterMessage(update *tgbotapi.Update, botAPI entity.BotAPI) {
	message := ""
	if update.Message.ForwardFrom != nil {
		message = strconv.Itoa(update.Message.ForwardFrom.ID)
	} else {
		message = strings.Replace(update.Message.Text, "<bq>", "`", -1)
		update.Message.Chat.ID = config.SquadID
	}
	msg, _ := botAPI.Send(entity.NewMessage(update, message))
	log.Println(msg.Text)
}
