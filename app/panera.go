package app

import (
	"github.com/andrysds/panera/app/bot"
	"github.com/andrysds/panera/app/web"
	"github.com/andrysds/panera/config"
	"gopkg.in/telegram-bot-api.v4"
)

type Panera interface {
	Run()
	SendMessage(*tgbotapi.MessageConfig)
}

func NewPanera() Panera {
	if config.App == "Bot" {
		return bot.NewBot()
	} else {
		return web.NewWeb()
	}
}
