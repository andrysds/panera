package app

import (
	"log"

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

func (p *Panera) Run() {
	log.Println("Panera starting...")

	started := make(chan bool, 2)
	go p.Web.Run(started)
	go p.Bot.Run(started)

	for i := 0; i < 2; i++ {
		<-started
	}
	p.IDle()
}

func (p *Panera) IDle() {
	log.Println("Use Ctrl-C to stop")
	for {
	}
}
