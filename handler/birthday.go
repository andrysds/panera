package handler

import (
	"fmt"
	"time"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleBirthdays(update *tgbotapi.Update, botAPI entity.BotAPI) {
	message := "Birthdays:"

	birthdays, err := entity.Birthdays(0, 0)
	clarity.PrintIfError(err, "error on getting birthdays")

	if err == nil {
		for _, b := range birthdays {
			message += fmt.Sprintf("\n%v %s - %s", b.Day, b.Month, b.Name)
		}
	}
	botAPI.Send(entity.NewMessage(update, message))
}

func HandleBirthdayKick(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if !IsFromMaster(update) {
		return
	}

	message := "birthday\\_kick"
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
				message += "\n" + b.Name + " gagal ditendang!"
			}
		}
		message += "\ndone"
	}
	botAPI.Send(entity.NewMessage(update, message))
}

func HandleBirthdayLink(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if !IsFromMaster(update) {
		return
	}
	message, err := botAPI.GetInviteLink(
		tgbotapi.ChatConfig{
			ChatID: config.BirthdayID,
		},
	)
	if err != nil {
		message = err.Error()
	}
	botAPI.Send(entity.NewMessage(update, message))
}

func HandleBirthdayUnban(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if !IsFromMaster(update) {
		return
	}
	_, err := botAPI.UnbanChatMember(
		tgbotapi.ChatMemberConfig{
			ChatID: config.BirthdayID,
			UserID: int(config.MasterID),
		},
	)
	message := ""
	if err != nil {
		message = "fail"
	} else {
		message = "success"
	}
	botAPI.Send(entity.NewMessage(update, message))
}
