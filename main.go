package main

import (
	"log"

	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/cron"
	"github.com/andrysds/panera/db"
)

func main() {
	log.Println("Panera starting...")
	config.Init()
	db.InitRedis()
	app.Init()
	cron.Init()
	app.Panera.Run()
}
