package cron

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
)

func AddStandupJobs() {
	Cron.AddFunc("0 27 13 * * 1-5", StandupJob)
	Cron.AddFunc("0 0 8 * * 1-5", StandupNewDayJob)
}

func StandupJob() {
	message := handler.HandleStandup(config.SquadID)
	app.Panera.SendMessage(message)
}

func StandupNewDayJob() {
	message := handler.HandleMasterCommand("standup_new_day", app.Panera.(*app.Bot).API)
	app.Panera.SendMessage(message)
}
