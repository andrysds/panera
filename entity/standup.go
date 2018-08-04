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
	Order    int
}

func NewStandup(data string, order int) *Standup {
	res := strings.Split(data, ":")
	standup := &Standup{}
	if len(res) == 3 {
		standup.Name = res[0]
		standup.Username = res[1]
		standup.State = res[2]
		standup.Order = order
	}
	return standup
}

func GetStandup() (*Standup, error) {
	order, err := db.Redis.Get(StandupKey).Int64()
	if err != nil {
		return nil, err
	}

	data, err := db.Redis.LRange(StandupListKey, order, order).Result()
	standup := &Standup{}
	if len(data) == 1 {
		standup = NewStandup(data[0], int(order))
	}
	return standup, err
}

func SkipStandup(standups []*Standup, currentOrder int) (*Standup, error) {
	for i := currentOrder + 1; i < len(standups); i++ {
		if standups[i].State != "1" {
			_, err := db.Redis.Set(StandupKey, i, 0).Result()
			return standups[i], err
		}
	}
	return nil, nil
}
