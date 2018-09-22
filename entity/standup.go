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
	obj := &Standup{}
	if len(res) == 3 {
		obj.Name = res[0]
		obj.Username = res[1]
		obj.State = res[2]
		obj.Order = order
	}
	return obj
}

func (self *Standup) Raw() string {
	return fmt.Sprintf("%s:%s:%s", self.Name, self.Username, self.State)
}

func CurrentStandup() (*Standup, error) {
	obj := &Standup{}
	order, err := db.Redis.Get(StandupKey).Int64()
	if err != nil {
		return nil, err
	}

	data, err := db.Redis.LRange(StandupListKey, order, order).Result()
	if len(data) == 1 {
		obj = NewStandup(data[0], int(order))
	}
	return obj, err
}

func NextStandup(fromBeginning bool) (*Standup, *Standup, error) {
	obj := &Standup{}
	current, err := CurrentStandup()
	if err != nil {
		return obj, current, err
	}

	objs, err := CurrentStandupList()
	if err != nil {
		return obj, current, err
	}

	i := current.Order + 1
	if fromBeginning {
		i = 0
	}

	for ; i < len(objs); i++ {
		if objs[i].State != "1" {
			_, err := db.Redis.Set(StandupKey, i, 0).Result()
			return objs[i], current, err
		}
	}
	return obj, current, errors.New("not found")
}

func NewDayStandup() string {
	current, err := CurrentStandup()
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

	if _, _, err := NextStandup(true); err != nil {
		return err.Error()
	}
	return "standup\\_new\\_day success"
}
