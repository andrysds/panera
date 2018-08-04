package migrate

import (
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/entity"
)

func StandupInit() string {
	if _, err := db.Redis.Del(entity.StandupKey).Result(); err != nil {
		return err.Error()
	}

	if _, err := db.Redis.Set(entity.StandupKey, 6, 0).Result(); err != nil {
		return err.Error()
	}

	return "standup_init success"
}
