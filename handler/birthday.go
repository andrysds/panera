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

	if err == nil {
		for _, b := range birthdays {
			messageText += fmt.Sprintf("\n%v %s - %s", b.Day, b.Month, b.Name)
		}
	}
	return entity.NewMessage(chatID, messageText)
}

func HandleBirthdayKick(botAPI *tgbotapi.BotAPI) string {
	messageText := "birthday\\_kick"
	tomorrow := time.Now().Add(24 * time.Hour)
	day := tomorrow.Day()
	month := tomorrow.Month()

	birthdays, err := entity.Birthdays(day, month)
	clarity.PrintIfError(err, "error on getting birthdays")

	if err == nil {
		for _, b := range birthdays {
			_, err := botAPI.KickChatMember(tgbotapi.KickChatMemberConfig{
				ChatMemberConfig: tgbotapi.ChatMemberConfig{
					ChatID: config.BirthdayID,
					UserID: b.UserID,
				},
			})
			if err != nil {
				messageText += "\n" + b.Name + " gagal ditendang!"
			}
		}
		messageText += "\ndone"
	}
	return messageText
}

func HandleBirthdayUnban(botAPI *tgbotapi.BotAPI) string {
	messageText, err := botAPI.GetInviteLink(
		tgbotapi.ChatConfig{
			ChatID: config.BirthdayID,
		},
	)
	if err != nil {
		messageText = err.Error()
	}
	return messageText
}
