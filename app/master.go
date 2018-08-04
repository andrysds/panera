package app

import (
	"github.com/andrysds/panera/db/migrate"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleMasterMessage(update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(p.ChatId, update.Message.Text)
	p.SendMessage(message)
}

func (p *Panera) HandleMasterCommand(command string) {
	result := ""
	switch command {
	case "standup_init":
		result = migrate.StandupInit()
	case "standup_list_init":
		result = migrate.StandupListInit()
	}
	message := tgbotapi.NewMessage(p.MasterId, result)
	p.SendMessage(message)
}
