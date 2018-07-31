package app

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) LogMessage(message *tgbotapi.Message) {
	fmt.Println("[%s] %s", message.From.UserName, message.Text)
}

func (p *Panera) SendMessage(message tgbotapi.MessageConfig) {
	_, err := p.Bot.Send(message)
	fmt.Println(err)
}
