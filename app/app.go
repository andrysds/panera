package app

import (
	"log"

	"github.com/andrysds/panera/app/bot"
	"github.com/andrysds/panera/app/web"
	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

var Panera App

type App interface {
	Run()
	SendMessage(*tgbotapi.MessageConfig)
}

func Init() {
	if config.App == "Bot" {
		Panera = bot.NewBot()
	} else {
		Panera = web.NewWeb()
	}
	log.Println("* App initialized")
}
