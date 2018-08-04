package app

import (
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) Run() {
	for update := range p.Updates {
		if update.Message == nil {
			continue
		}
		p.LogMessage(update.Message)

		var message *tgbotapi.MessageConfig
		switch {
		case update.Message.NewChatMembers != nil:
			message = p.HandleGroupInvitation(&update)
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "standup":
				message = p.HandleStandup(&update)
			case "standup_list":
				message = p.HandleStandupList(&update)
			case "standup_skip":
				message = p.HandleStandupSkip(&update)
			default:
				if update.Message.Chat.ID == p.MasterId {
					message = p.HandleMasterCommand(update.Message.Command())
				}
			}
		case update.Message.Chat.ID == p.MasterId:
			p.HandleMasterMessage(&update)
		}
		p.SendMessage(message)
	}
}
