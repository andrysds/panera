package cron

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
)

func AddBirthdayJobs() {
	Cron.AddFunc("0 0 8 * * *", BirthdayKickJob)
}

func BirthdayKickJob() {
	handler.HandleBirthdayKick(
		UpdateFromCron(config.MasterID),
		app.Panera.GetBotAPI(),
	)
}
