package main

import (
	"log"
	"net/http"

	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/route"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Panera starting...")

	gotenv.Load()
	db.Init()
	route.Init()

	log.Println("* Listening on :8080")
	log.Println("Use Ctrl-C to stop")
	http.ListenAndServe(":8080", nil)
}
