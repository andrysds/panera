package entity

import (
	"errors"
	"fmt"
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

func (self *Standup) Raw() string {
	return fmt.Sprintf("%s:%s:%s", self.Name, self.Username, self.State)
}

func StandupCurrent() (*Standup, error) {
	standup := &Standup{}
	order, err := db.Redis.Get(StandupKey).Int64()
	if err != nil {
		return nil, err
	}

	data, err := db.Redis.LRange(StandupListKey, order, order).Result()
	if len(data) == 1 {
		standup = NewStandup(data[0], int(order))
	}
	return standup, err
}

func StandupNext() (*Standup, *Standup, error) {
	standup := &Standup{}
	current, err := StandupCurrent()
	if err != nil {
		return standup, current, err
	}

	standups, err := StandupListCurrent()
	if err != nil {
		return standup, current, err
	}

	for i := current.Order + 1; i < len(standups); i++ {
		if standups[i].State != "1" {
			_, err := db.Redis.Set(StandupKey, i, 0).Result()
			return standups[i], current, err
		}
	}
	return standup, current, errors.New("not found")
}

func StandupNewDay() string {
	current, err := StandupCurrent()
	if err != nil {
		return err.Error()
	}

	current.State = "1"
	if _, err := db.Redis.LSet(StandupListKey, int64(current.Order), current.Raw()).Result(); err != nil {
		return err.Error()
	}

	if _, err := db.Redis.Set(StandupKey, 0, 0).Result(); err != nil {
		return err.Error()
	}

	if _, _, err := StandupNext(); err != nil {
		return err.Error()
	}
	return "standup_new_day success"
}
