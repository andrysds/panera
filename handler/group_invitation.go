package handler

import (
	"strconv"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func IsAddedToGroup(update *tgbotapi.Update) bool {
	if update.Message.NewChatMembers != nil {
		members := update.Message.NewChatMembers
		for _, member := range *members {
			if member.UserName == "panera_bot" {
				return true
			}
		}
	}
	return false
}

func HandleGroupInvitation(update *tgbotapi.Update, botAPI entity.BotAPI) {
	message := "I was invited to " + strconv.FormatInt(update.Message.Chat.ID, 10)
	update.Message.Chat.ID = config.MasterID
	update.Message.MessageID = 0
	botAPI.Send(entity.NewMessage(update, message))
}
