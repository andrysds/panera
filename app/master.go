package app

import (
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleMasterMessage(update *tgbotapi.Update) {
	message := tgbotapi.NewMessage(p.ChatId, update.Message.Text)
	p.SendMessage(message)
}

func (p *Panera) HandleMasterCommand(command string) {
	result := "command is not defined"
	switch command {
	// migrate
	case "standup_init":
		result = migrate.StandupInit()
	case "standup_list_init":
		result = migrate.StandupListInit()

	// standup
	case "standup_new_day":
		result = entity.StandupNewDay()
	}
	message := tgbotapi.NewMessage(p.MasterId, result)
	p.SendMessage(message)
}
