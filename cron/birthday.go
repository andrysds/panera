package cron

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/handler"
)

func AddBirthdayJobs() {
	Cron.AddFunc("0 0 8 * * *", BirthdayKickJob)
}

func BirthdayKickJob() {
	message := handler.HandleMasterCommand("birthday_kick", app.Panera.BotAPI())
	app.Panera.SendMessage(message)
}
