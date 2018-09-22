package app

import (
	"log"

	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

var Panera App

type App interface {
	Run()
	SendMessage(*tgbotapi.MessageConfig)
}

func Init() {
	if config.BotToken != "" {
		Panera = NewBot()
	} else {
		Panera = NewCli()
	}
	log.Println("* App initialized")
}
