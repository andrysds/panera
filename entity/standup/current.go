package standup

import (
	"github.com/andrysds/panera/db"
)

func Current() (*Standup, error) {
	obj := &Standup{}
	order, err := db.Redis.Get(Key).Int64()
	if err != nil {
		return nil, err
	}

	data, err := db.Redis.LRange(ListKey, order, order).Result()
	if len(data) == 1 {
		obj = NewStandup(data[0], int(order))
	}
	return obj, err
}
