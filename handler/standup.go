package handler

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func HandleStandup(update *tgbotapi.Update, botAPI entity.BotAPI) {
	standup, err := entity.CurrentStandup()
	clarity.PrintIfError(err, "error on get standup")

	messageTemplate := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah _%s_ (@%s)"
	message := fmt.Sprintf(messageTemplate, standup.Name, standup.Username)
	botAPI.Send(entity.NewMessage(update, message))
}

func HandleStandupList(update *tgbotapi.Update, botAPI entity.BotAPI) {
	standups, err := entity.CurrentStandupList()
	clarity.PrintIfError(err, "error on get standup list")

	message := "Stand up lead periode ini:"
	for _, s := range standups {
		message += "\n"
		if s.State == "1" {
			message += "`[x]` "
		} else {
			message += "`[ ]` "
		}
		message += s.Name
	}
	botAPI.Send(entity.NewMessage(update, message))
}

func HandleStandupNewDay(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if !IsFromMaster(update) {
		return
	}

	result := entity.NewDayStandup()
	if result == entity.NotFoundMessage {
		HandleStandupNewPeriod(update, botAPI)
	} else {
		botAPI.Send(entity.NewMessage(update, result))
	}
}

func HandleStandupNewPeriod(update *tgbotapi.Update, botAPI entity.BotAPI) {
	if !IsFromMaster(update) {
		return
	}

	result := entity.NewPeriodStandupList()
	botAPI.Send(entity.NewMessage(update, result))

	update.Message.Chat.ID = config.SquadID
	HandleStandupList(update, botAPI)
}

func HandleStandupSkip(update *tgbotapi.Update, botAPI entity.BotAPI) {
	standup, current, err := entity.NextStandup(false)
	clarity.PrintIfError(err, "error on skipping standup")

	message := ""
	if err == nil {
		message = "Karena %s tidak bisa, penggantinya _%s_ (@%s)"
		message = fmt.Sprintf(message, current.Name, standup.Name, standup.Username)
	} else if err.Error() == entity.NotFoundMessage {
		message = "Waduh ga ada gantinya lagi nih!"
	}
	botAPI.Send(entity.NewMessage(update, message))
}
