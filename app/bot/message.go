package bot

import (
	"log"

	"github.com/andrysds/clarity"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		_, err := b.API.Send(message)
		clarity.PrintIfError(err, "error on sending message")
	}
}

func (b *Bot) LogMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s\n", message.From.UserName, message.Text)
}
