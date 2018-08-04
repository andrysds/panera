package main

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/db"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	db.InitRedis()
	defer db.Redis.Close()

	panera := app.NewPanera()
	panera.Run()
}
