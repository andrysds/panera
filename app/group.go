package app

import (
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleGroupInvitation(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	if !p.IsAddedToGroup(update.Message.NewChatMembers) {
		return nil
	}
	messageText := "I was invited to " + strconv.FormatInt(update.Message.Chat.ID, 10)
	message := p.NewMessage(p.MasterId, messageText)
	return message
}

func (p *Panera) IsAddedToGroup(members *[]tgbotapi.User) bool {
	for _, member := range *members {
		if member.UserName == "panera_bot" {
			return true
		}
	}
	return false
}
