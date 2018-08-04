package main

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db"
)

func main() {
	config.Init()

	db.InitRedis()
	defer db.Redis.Close()

	panera := app.NewPanera()
	panera.Run()
}
