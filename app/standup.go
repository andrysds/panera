package app

import (
	"fmt"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/entity"
	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleStandup(update *tgbotapi.Update) {
	standup, err := entity.GetStandup()
	clarity.PrintIfError(err, "error on get standup")

	messageTemplate := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah _%s_ (@%s)"
	messageText := fmt.Sprintf(messageTemplate, standup.Name, standup.Username)

	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

func (p *Panera) HandleStandupList(update *tgbotapi.Update) {
	standups, err := entity.GetStandupList()
	clarity.PrintIfError(err, "error on get standup list")

	messageText := "Stand up lead periode ini:\n"
	for _, s := range standups {
		if s.State == "1" {
			messageText += "`[x]` "
		} else {
			messageText += "`[ ]` "
		}
		messageText += s.Name + "\n"
	}

	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}
