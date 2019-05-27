package router

import (
	"github.com/andrysds/panera/connection"
	"github.com/andrysds/panera/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TelegramRouter is router function for telegram updates
func TelegramRouter() {
	if connection.Telegram != nil {
		updates := connection.Telegram.ListenForWebhook("/" + connection.Telegram.Token)
		for u := range updates {
			if u.Message == nil {
				return
			}
			switch {
			case isAddedToGroup(&u):
				handler.GroupInvitationMessage(&u)
			case u.Message.IsCommand():
				handler.CommandMessage(u.Message.Command())
			default:
				handler.MasterMessage(&u)
			}
		}
	}
}

func isAddedToGroup(update *tgbotapi.Update) bool {
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
