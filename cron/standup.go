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
	handler.HandleStandup(
		UpdateFromCron(config.SquadID),
		app.Panera.GetBotAPI(),
	)
}

func StandupNewDayJob() {
	handler.HandleStandupNewDay(
		UpdateFromCron(config.MasterID),
		app.Panera.GetBotAPI(),
	)
}
