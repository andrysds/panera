package main

import (
	"os"
	"os/signal"

	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/scheduler/job"
	"github.com/robfig/cron"
)

type Scheduler struct {
	Cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Cron: cron.New(),
	}
}

func (s *Scheduler) AddFuncs() {
	s.Cron.AddFunc("0 27 13 * * 1-5", job.Standup)
	s.Cron.AddFunc("0 0 8 * * 1-5", job.StandupNewDay)
}

func (s *Scheduler) Start() {
	go s.Cron.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func main() {
	config.Init()
	db.InitRedis()
	app.Init()

	scheduler := NewScheduler()
	scheduler.AddFuncs()
	scheduler.Start()
}
