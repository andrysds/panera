package app

import "gopkg.in/telegram-bot-api.v4"

func (p *Panera) HandleStandup(update *tgbotapi.Update) {
	messageText := "Yuk stand up! Yang dapat giliran untuk memimpin stand up hari ini adalah <b>Yohanes</b>"
	message := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	p.SendMessage(message)
}
