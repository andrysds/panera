package app

import (
	"github.com/andrysds/panera/app/bot"
	"github.com/andrysds/panera/app/web"
	"github.com/andrysds/panera/config"
)

type Panera interface {
	Run()
}

func NewPanera() Panera {
	if config.App == "Bot" {
		return bot.NewBot()
	} else {
		return web.NewWeb()
	}
}
