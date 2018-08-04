package bot

import (
	"fmt"

	"github.com/andrysds/clarity"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) NewMessage(chatId int64, text string) *tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(chatId, text)
	message.ParseMode = "markdown"
	return &message
}

func (b *Bot) LogMessage(message *tgbotapi.Message) {
	fmt.Printf("[%s] %s\n", message.From.UserName, message.Text)
}

func (b *Bot) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		_, err := b.BotAPI.Send(message)
		clarity.PrintIfError(err, "error on sending message")
	}
}
