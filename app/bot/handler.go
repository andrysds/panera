package bot

import (
	"strconv"

	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	message := handler.NewMessage(b.ChatId, update.Message.Text)
	return message
}

func (b *Bot) HandleGroupInvitation(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	if !b.IsAddedToGroup(update.Message.NewChatMembers) {
		return nil
	}
	messageText := "I was invited to " + strconv.FormatInt(update.Message.Chat.ID, 10)
	message := handler.NewMessage(b.MasterId, messageText)
	return message
}

func (b *Bot) IsAddedToGroup(members *[]tgbotapi.User) bool {
	for _, member := range *members {
		if member.UserName == "panera_bot" {
			return true
		}
	}
	return false
}
