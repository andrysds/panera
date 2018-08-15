package main

import (
	"log"

	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db"
)

func main() {
	log.Println("Panera starting...")

	config.Init()
	db.InitRedis()
	app.Init()

	log.Println("* Listening on tcp://0.0.0.0:" + config.Port)
	log.Println("Use Ctrl-C to stop")

	app.Panera.Run()
}
