package main

import (
	"net/http"
	"os"

	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/db"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	panera := app.NewPanera()

	db.InitRedis()
	defer db.Redis.Close()

	port := os.Getenv("PORT")
	go http.ListenAndServe(":"+port, nil)

	panera.Run()
}
