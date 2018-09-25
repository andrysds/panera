package entity

import (
	"math/rand"
	"time"

	"github.com/andrysds/panera/db"
)

const StandupListKey = "panera:standup:list"

type StandupList []*Standup

func CurrentStandupList() (StandupList, error) {
	result, err := db.Redis.LRange(StandupListKey, 0, -1).Result()
	obj := StandupList{}
	for i, r := range result {
		obj = append(obj, NewStandup(r, i))
	}
	return obj, err
}

func NewPeriodStandupList() string {
	currentList, err := CurrentStandupList()
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
		if _, err := db.Redis.LSet(StandupListKey, int64(i), standup.Raw()).Result(); err != nil {
			return err.Error()
		}
	}

	if _, err := db.Redis.Set(StandupKey, 0, 0).Result(); err != nil {
		return err.Error()
	}

	return "standup\\_new\\_period\ndone"
}
