package helper

import (
	"io/ioutil"
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckAndFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
}

func SendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig) {
	_, err := bot.Send(message)
	Check(err)
}
