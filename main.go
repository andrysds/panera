package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andrysds/panera/connection"
	"github.com/andrysds/panera/cron"
	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/router"
	"github.com/andrysds/panera/template"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Panera starting...")

	gotenv.Load()
	connection.Init()
	entity.InitCollection()
	cron.Init()
	template.Init()

	port := os.Getenv("PORT")
	log.Println("* Listening on :" + port)
	log.Println("Use Ctrl-C to stop")
	http.ListenAndServe(":"+port, router.Router())
}
