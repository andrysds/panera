package handler

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleBirthdays(chatID int64) *tgbotapi.MessageConfig {
	birthdays, err := entity.Birthdays(0, 0)
	clarity.PrintIfError(err, "error on get birthdays")

	messageText := "Birthdays:"
	for _, b := range birthdays {
		messageText += fmt.Sprintf("\n%v %s - %s", b.Day, b.Month, b.Name)
	}

	message := entity.NewMessage(chatID, messageText)
	return message
}
