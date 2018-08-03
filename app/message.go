package app

import (
	"fmt"

	"github.com/andrysds/clarity"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) NewMessage(chatId int64, text string) tgbotapi.MessageConfig {
	message := tgbotapi.NewMessage(chatId, text)
	message.ParseMode = "markdown"
	return message
}

func (p *Panera) LogMessage(message *tgbotapi.Message) {
	fmt.Println("[%s] %s", message.From.UserName, message.Text)
}

func (p *Panera) SendMessage(message tgbotapi.MessageConfig) {
	_, err := p.Bot.Send(message)
	clarity.PrintIfError(err, "error on sending message")
}
