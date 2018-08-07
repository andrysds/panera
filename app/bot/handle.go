package bot

import (
	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) Handle(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	var message *tgbotapi.MessageConfig
	chatID := update.Message.Chat.ID
	handler.LogMessage(update.Message.From.UserName, update.Message.Text)

	switch {
	case b.IsAddedToGroup(update):
		message = handler.HandleGroupInvitation(chatID)
	case update.Message.IsCommand():
		message = handler.HandleCommand(chatID, update.Message.Command())
	case chatID == config.MasterID:
		message = handler.HandleMasterMessage(update)
	}
	b.SendMessage(message)
}

func (b *Bot) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		_, err := b.API.Send(message)
		clarity.PrintIfError(err, "error on sending message")
	}
}

func (b *Bot) IsAddedToGroup(update *tgbotapi.Update) bool {
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
