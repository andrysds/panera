package app

import (
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleGroupInvitation(update *tgbotapi.Update) {
	if !p.IsAddedToGroup(update.Message.NewChatMembers) {
		return
	}
	messageText := "I was invited to " + strconv.FormatInt(p.ChatId, 10)
	message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

func (p *Panera) IsAddedToGroup(members *[]tgbotapi.User) bool {
	for _, member := range *members {
		if member.UserName == "panera_bot" {
			return true
		}
	}
	return false
}
