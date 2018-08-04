package bot

import (
	"log"

	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) Run(started chan<- bool) {
	b.BotAPI = NewBotAPI()
	b.Updates = NewUpdates(b.BotAPI)
	log.Println("* [bot] Listening from webhook")
	started <- true

	for update := range b.Updates {
		if update.Message == nil {
			continue
		}
		b.LogMessage(update.Message)

		var message *tgbotapi.MessageConfig
		chatId := update.Message.Chat.ID

		switch {
		case update.Message.NewChatMembers != nil:
			message = b.HandleGroupInvitation(&update)
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "standup":
				message = handler.HandleStandup(chatId)
			case "standup_list":
				message = handler.HandleStandupList(chatId)
			case "standup_skip":
				message = handler.HandleStandupSkip(chatId)
			default:
				if chatId == b.MasterId {
					message = handler.HandleMasterCommand(chatId, update.Message.Command())
				}
			}
		case chatId == b.MasterId:
			b.HandleMasterMessage(&update)
		}
		b.SendMessage(message)
	}
}
