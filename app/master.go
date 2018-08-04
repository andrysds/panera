package app

import (
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	message := p.NewMessage(p.ChatId, update.Message.Text)
	return message
}

func (p *Panera) HandleMasterCommand(command string) *tgbotapi.MessageConfig {
	result := "command is not defined"
	switch command {
	// migrate
	case "standup_init":
		result = migrate.StandupInit()
	case "standup_list_init":
		result = migrate.StandupListInit()

	// standup
	case "standup_new_day":
		result = standup.NewDay()
	}
	message := p.NewMessage(p.MasterId, result)
	return message
}
