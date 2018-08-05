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

	panera := app.NewPanera()
	log.Println("Use Ctrl-C to stop")

	panera.Run()
}
