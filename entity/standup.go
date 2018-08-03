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
	standup, err := db.Redis.Get(StandupKey).Result()
	return NewStandup(standup), err
}
