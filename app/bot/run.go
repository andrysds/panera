package bot

import (
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) Run() {
	for update := range b.Updates {
		if update.Message == nil {
			continue
		}
		b.LogMessage(update.Message)

		var message *tgbotapi.MessageConfig
		switch {
		case update.Message.NewChatMembers != nil:
			message = b.HandleGroupInvitation(&update)
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "standup":
				message = b.HandleStandup(&update)
			case "standup_list":
				message = b.HandleStandupList(&update)
			case "standup_skip":
				message = b.HandleStandupSkip(&update)
			default:
				if update.Message.Chat.ID == b.MasterId {
					message = b.HandleMasterCommand(update.Message.Command())
				}
			}
		case update.Message.Chat.ID == b.MasterId:
			b.HandleMasterMessage(&update)
		}
		b.SendMessage(message)
	}
}
