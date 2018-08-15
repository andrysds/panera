package job

import (
	"github.com/andrysds/panera/app"
	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
)

func Standup() {
	message := handler.HandleStandup(config.SquadID)
	app.Panera.SendMessage(message)
}

func StandupNewDay() {
	message := handler.HandleMasterCommand("standup_new_day")
	app.Panera.SendMessage(message)
}
