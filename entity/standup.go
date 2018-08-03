package entity

import (
	"strings"

	"github.com/andrysds/panera/db"
)

const StandupKey = "panera:standup"

type Standup struct {
	Name     string
	Username string
	State    string
}

func NewStandup(data string) *Standup {
	res := strings.Split(data, ":")
	standup := &Standup{}
	if len(res) == 3 {
		standup.Name = res[0]
		standup.Username = res[1]
		standup.State = res[2]
	}
	return standup
}

func GetStandup() (*Standup, error) {
	current, err := db.Redis.Get(StandupKey).Int64()
	if err != nil {
		return nil, err
	}
	data, err := db.Redis.LRange(StandupListKey, current, current).Result()
	standup := &Standup{}
	if len(data) == 1 {
		standup = NewStandup(data[0])
	}
	return standup, err
}
