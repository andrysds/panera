package standup

import (
	"math/rand"
	"time"

	"github.com/andrysds/panera/db"
)

func NewPeriod() string {
	currentList, err := CurrentList()
	if err != nil {
		return err.Error()
	}

	rand.Seed(time.Now().UnixNano())
	for i := range currentList {
		j := rand.Intn(i + 1)
		currentList[i], currentList[j] = currentList[j], currentList[i]
	}

	for i, standup := range currentList {
		standup.State = "0"
		if _, err := db.Redis.LSet(ListKey, int64(i), standup.Raw()).Result(); err != nil {
			return err.Error()
		}
	}

	if _, err := db.Redis.Set(Key, 0, 0).Result(); err != nil {
		return err.Error()
	}

	return "standup\\_new\\_period success"
}
