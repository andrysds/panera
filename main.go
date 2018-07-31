package main

import (
	"net/http"
	"os"

	"github.com/andrysds/pnr_bot/app"
)

func main() {
	port := os.Getenv("PORT")
	go http.ListenAndServe(":"+port, nil)

	pnr_bot := app.NewPnrBot()
	pnr_bot.Run()
}
