package standup

import (
	"github.com/andrysds/panera/db"
)

func CurrentList() (StandupList, error) {
	result, err := db.Redis.LRange(ListKey, 0, -1).Result()
	obj := StandupList{}
	for i, r := range result {
		obj = append(obj, NewStandup(r, i))
	}
	return obj, err
}
