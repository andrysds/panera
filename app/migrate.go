package app

import (
	"github.com/andrysds/panera/db/migrate"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleInitStandup(update *tgbotapi.Update) {
	result := migrate.InitStandup()
	message := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	p.SendMessage(message)
}

func (p *Panera) HandleInitStandupList(update *tgbotapi.Update) {
	result := migrate.InitStandupList()
	message := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	p.SendMessage(message)
}
