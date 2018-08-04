package entity

import (
	"github.com/andrysds/panera/db"
)

const StandupListKey = "panera:standup:list"

type StandupList []*Standup

func GetStandupList() (StandupList, error) {
	result, err := db.Redis.LRange(StandupListKey, 0, -1).Result()
	standupList := []*Standup{}
	for i, r := range result {
		standupList = append(standupList, NewStandup(r, i))
	}
	return standupList, err
}
