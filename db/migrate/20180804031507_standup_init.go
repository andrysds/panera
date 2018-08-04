package migrate

import (
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/entity/standup"
)

func StandupInit() string {
	if _, err := db.Redis.Del(standup.Key).Result(); err != nil {
		return err.Error()
	}

	if _, err := db.Redis.Set(standup.Key, 6, 0).Result(); err != nil {
		return err.Error()
	}

	return SuccessMessage
}
