package app

import (
	"fmt"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleStandup(update *tgbotapi.Update) {
	messageText := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah "
	for _, data := range dataExamples {
		person := strings.Split(data, ":")
		if person[2] == "0" {
			messageText += fmt.Sprintf("_%s_ (@%s)", person[0], person[1])
			break
		}
	}
	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

func (p *Panera) HandleStandupList(update *tgbotapi.Update) {
	messageText := "Stand up lead periode ini:\n"
	for _, data := range dataExamples {
		person := strings.Split(data, ":")
		if person[2] == "1" {
			messageText += "`[x]` "
		} else {
			messageText += "`[ ]` "
		}
		messageText += person[0] + "\n"
	}

	message := p.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

var dataExamples = []string{
	"Herry:herrydev:1",
	"Olvi:olvilora:1",
	"Isti:tianaulia:1",
	"Setia:setiasimaremare:1",
	"Rifa:rifaMukhlisa:1",
	"Yohanes:yohanes77:0",
	"Regina:regina_avena:0",
	"Ben:benlemueltanasale:0",
	"Farida:faridaamila:0",
	"Adimas:addimas:0",
	"Andrys:andrysds:0",
	"Ai:ayshzkh:0",
}
