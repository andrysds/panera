package handler

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleStandup(chatID int64) *tgbotapi.MessageConfig {
	standup, err := entity.CurrentStandup()
	clarity.PrintIfError(err, "error on get standup")

	messageTemplate := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah _%s_ (@%s)"
	messageText := fmt.Sprintf(messageTemplate, standup.Name, standup.Username)

	message := entity.NewMessage(chatID, messageText)
	return message
}

func HandleStandupSkip(chatID int64) *tgbotapi.MessageConfig {
	standup, current, err := entity.NextStandup(false)
	clarity.PrintIfError(err, "error on skipping standup")

	if err == nil {
		messageTemplate := "Karena %s tidak bisa, penggantinya _%s_ (@%s)"
		messageText := fmt.Sprintf(messageTemplate, current.Name, standup.Name, standup.Username)
		message := entity.NewMessage(chatID, messageText)
		return message
	} else if err.Error() == "not found" {
		messageText := "Waduh ga ada gantinya lagi nih!"
		message := entity.NewMessage(chatID, messageText)
		return message
	}
	return nil
}
