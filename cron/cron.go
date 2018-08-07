package cron

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func Init() {
	Cron := cron.New()
	Cron.AddFunc("0 30 * * * *", TestCron)
	go Cron.Start()
	log.Println("* Cron initialized")
}

func TestCron() {
	log.Println("Every hour on the half hour", time.Now())
}
