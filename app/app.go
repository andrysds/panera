package app

import (
	"log"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/entity"
)

var Panera App

type App interface {
	GetBotAPI() entity.BotAPI
	Run()
}

func Init() {
	if config.BotToken != "" {
		Panera = NewBot()
	} else {
		Panera = NewCli()
	}
	log.Println("* App initialized")
}
