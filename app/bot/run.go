package bot

import (
	"log"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) Run(started chan<- bool) {
	log.Println("* [bot] Listening from webhook")
	started <- true

	for update := range b.Updates {
		if update.Message == nil {
			continue
		}
		b.LogMessage(update.Message)

		var message *tgbotapi.MessageConfig
		chatID := update.Message.Chat.ID

		switch {
		case b.IsAddedToGroup(&update):
			message = handler.HandleGroupInvitation(chatID)
		case update.Message.IsCommand():
			command := update.Message.Command()
			switch command {
			case "standup":
				message = handler.HandleStandup(chatID)
			case "standup_list":
				message = handler.HandleStandupList(chatID)
			case "standup_skip":
				message = handler.HandleStandupSkip(chatID)
			default:
				if chatID == config.MasterID {
					message = handler.HandleMasterCommand(command)
				}
			}
		case chatID == config.MasterID:
			handler.HandleMasterMessage(&update)
		}
		b.SendMessage(message)
	}
}
