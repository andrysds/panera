package bot

import (
	"github.com/andrysds/panera/db/migrate"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) HandleMasterMessage(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	message := b.NewMessage(b.ChatId, update.Message.Text)
	return message
}

func (b *Bot) HandleMasterCommand(command string) *tgbotapi.MessageConfig {
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
	message := b.NewMessage(b.MasterId, result)
	return message
}
