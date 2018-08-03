package migrate

import (
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/entity"
)

func InitStandup() string {
	_, err := db.Redis.Del(entity.StandupKey).Result()
	if err != nil {
		return err.Error()
	}

	data := "Regina:regina\\_avena:0"
	_, err = db.Redis.Set(entity.StandupKey, data, 0).Result()
	if err != nil {
		return err.Error()
	}

	return "ok"
}
