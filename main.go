package main

import (
	"net/http"
	"os"

	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/db"
)

func main() {
	port := os.Getenv("PORT")
	go http.ListenAndServe(":"+port, nil)

	db.InitRedis()
	defer db.Redis.Close()

	panera := app.NewPanera()
	panera.Run()
}
