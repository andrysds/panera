package app

import (
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleMasterMessage(update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(p.ChatId, update.Message.Text)
	p.SendMessage(message)
}
