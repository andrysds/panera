package cron

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func Init() {
	Cron := cron.New()
	Cron.AddFunc("0 * * * * *", TestCron)
	go Cron.Start()
	log.Println("* Cron initialized")
}

func TestCron() {
	log.Println("Every minutes", time.Now())
}
