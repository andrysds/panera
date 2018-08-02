package app

import (
	"github.com/andrysds/panera/db/migrate"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleMasterMessage(update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(p.ChatId, update.Message.Text)
	p.SendMessage(message)
}

func (p *Panera) HandleInitStandup(update *tgbotapi.Update) {
	result := migrate.InitStandup()
	message := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	p.SendMessage(message)
}
