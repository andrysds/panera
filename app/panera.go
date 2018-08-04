package app

import (
	"github.com/andrysds/panera/app/bot"
	"github.com/andrysds/panera/app/web"
)

type Panera struct {
	Bot *bot.Bot
	Web *web.Web
}

func NewPanera() *Panera {
	return &Panera{
		Bot: bot.NewBot(),
		Web: web.NewWeb(),
	}
}
