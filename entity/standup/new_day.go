package standup

import (
	"github.com/andrysds/panera/db"
)

func NewDay() string {
	current, err := Current()
	if err != nil {
		return err.Error()
	}

	current.State = "1"
	if _, err := db.Redis.LSet(ListKey, int64(current.Order), current.Raw()).Result(); err != nil {
		return err.Error()
	}

	if _, err := db.Redis.Set(Key, 0, 0).Result(); err != nil {
		return err.Error()
	}

	if _, _, err := Next(true); err != nil {
		return err.Error()
	}
	return "standup\\_new\\_day success"
}
