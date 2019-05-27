package handler

import (
	"strconv"

	"github.com/andrysds/panera/connection"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CommandMessage handles telegram command messages
func CommandMessage(update *tgbotapi.Update) {
	cmd := update.Message.Command()
	if cmd == "standup_new_day" && update.Message.Chat.ID != connection.MasterTelegramID {
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, Command(cmd))
	connection.Telegram.Send(msg)
}

// GroupInvitationMessage handles when there is group invitation event
func GroupInvitationMessage(update *tgbotapi.Update) {
	msgTxt := "I was invited to " + strconv.FormatInt(update.Message.Chat.ID, 10)
	msg := tgbotapi.NewMessage(connection.MasterTelegramID, msgTxt)
	connection.Telegram.Send(msg)
}

// MasterMessage handles messages from master
func MasterMessage(update *tgbotapi.Update) {
	if update.Message.Chat.ID == connection.MasterTelegramID {
		var msg tgbotapi.MessageConfig
		if update.Message.ForwardFrom != nil {
			msg = tgbotapi.NewMessage(connection.MasterTelegramID, strconv.Itoa(update.Message.ForwardFrom.ID))
		} else {
			msg = tgbotapi.NewMessage(connection.SquadTelegramID, update.Message.Text)
		}
		connection.Telegram.Send(msg)
	}
}
