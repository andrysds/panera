package app

import (
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func (p *Panera) HandleStandup(update *tgbotapi.Update) {
	messageText := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah Yohanes"
	message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}

func (p *Panera) HandleStandupList(update *tgbotapi.Update) {
	messageText := ""
	for _, data := range dataExamples {
		person := strings.Split(data, ":")
		if person[2] == "1" {
			messageText += "`[x]`"
		} else {
			messageText += "`[ ]`"
		}
		messageText += person[0] + "\n"
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	message.ParseMode = "markdown"
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
