package main

import (
	"log"

	"github.com/andrysds/panera/db"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Panera starting...")
	gotenv.Load()
	db.Init()
	db.Test()
}
