package migrate

import (
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/entity"
)

func StandupInit() string {
	_, err := db.Redis.Del(entity.StandupKey).Result()
	if err != nil {
		return err.Error()
	}

	_, err = db.Redis.Set(entity.StandupKey, 6, 0).Result()
	if err != nil {
		return err.Error()
	}

	return "standup_init success"
}
