package handler

import (
	"fmt"
	"time"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleBirthdays(chatID int64) *tgbotapi.MessageConfig {
	messageText := "Birthdays:"

	birthdays, err := entity.Birthdays(0, 0)
	clarity.PrintIfError(err, "error on getting birthdays")

	if err != nil {
		return entity.NewMessage(chatID, messageText)
	} else {
		for _, b := range birthdays {
			messageText += fmt.Sprintf("\n%v %s - %s", b.Day, b.Month, b.Name)
		}
	}
	return entity.NewMessage(chatID, messageText)
}

func HandleBirthdayKick(botAPI *tgbotapi.BotAPI) string {
	result := "birthday\\_kick "
	now := time.Now()
	day := now.Day()
	month := now.Month()

	birthdays, err := entity.Birthdays(day, month)
	clarity.PrintIfError(err, "error on getting birthdays")

	if err != nil {
		result += "fail"
	} else {
		for _, b := range birthdays {
			_, err := botAPI.KickChatMember(tgbotapi.KickChatMemberConfig{
				ChatMemberConfig: tgbotapi.ChatMemberConfig{
					ChatID: config.BirthdayID,
					UserID: b.UserID,
				},
				UntilDate: now.Add(24 * time.Hour).Unix(),
			})
			clarity.PrintIfError(err, "error on kicking chat member")
		}

		result += "success"
	}
	return result
}
