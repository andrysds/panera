package bot

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity/standup"
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) HandleStandup(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	standup, err := standup.Current()
	clarity.PrintIfError(err, "error on get standup")

	messageTemplate := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah _%s_ (@%s)"
	messageText := fmt.Sprintf(messageTemplate, standup.Name, standup.Username)

	message := b.NewMessage(update.Message.Chat.ID, messageText)
	return message
}

func (b *Bot) HandleStandupSkip(update *tgbotapi.Update) *tgbotapi.MessageConfig {
	standup, current, err := standup.Next()
	clarity.PrintIfError(err, "error on skipping standup")

	if err == nil {
		messageTemplate := "Karena %s tidak bisa, penggantinya _%s_ (@%s)"
		messageText := fmt.Sprintf(messageTemplate, current.Name, standup.Name, standup.Username)
		message := b.NewMessage(update.Message.Chat.ID, messageText)
		return message
	} else if err.Error() == "not found" {
		messageText := "Waduh ga ada gantinya lagi nih!"
		message := b.NewMessage(update.Message.Chat.ID, messageText)
		return message
	}
	return nil
}
