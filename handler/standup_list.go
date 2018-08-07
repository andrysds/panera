package handler

import (
	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleStandupList(chatID int64) *tgbotapi.MessageConfig {
	standups, err := standup.CurrentList()
	clarity.PrintIfError(err, "error on get standup list")

	messageText := "Stand up lead periode ini:\n"
	for _, s := range standups {
		if s.State == "1" {
			messageText += "`[x]` "
		} else {
			messageText += "`[ ]` "
		}
		messageText += s.Name + "\n"
	}

	message := entity.NewMessage(chatID, messageText)
	return message
}
