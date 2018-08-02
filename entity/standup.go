package entity

import (
	"strings"

	"github.com/andrysds/panera/db"
)

const StandupKey = "panera:standup"

type Standup struct {
	Name     string
	Username string
	HasDone  bool
}

func NewStandup(data string) *Standup {
	standup := strings.Split(data, ":")

	hasDone := false
	if standup[2] == "1" {
		hasDone = true
	}

	return &Standup{
		Name:     standup[0],
		Username: standup[1],
		HasDone:  hasDone,
	}
}

func GetStandup() (*Standup, error) {
	standupList, err := GetStandupList()
	for _, s := range standupList {
		if !s.HasDone {
			return s, err
		}
	}
	return &Standup{}, err
}

func GetStandupList() ([]*Standup, error) {
	result, err := db.Redis.LRange(StandupKey, 0, -1).Result()
	standupList := []*Standup{}
	for _, r := range result {
		standupList = append(standupList, NewStandup(r))
	}
	return standupList, err
}
