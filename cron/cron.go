package cron

import (
	"log"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func Init() {
	Cron = cron.New()

	AddStandupJobs()

	go Cron.Start()
	log.Println("* Cron initialized")
}