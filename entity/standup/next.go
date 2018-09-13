package standup

import (
	"errors"

	"github.com/andrysds/panera/db"
)

func Next() (*Standup, *Standup, error) {
	obj := &Standup{}
	current, err := Current()
	if err != nil {
		return obj, current, err
	}

	objs, err := CurrentList()
	if err != nil {
		return obj, current, err
	}

	for _, obj := range objs {
		if obj.State != "1" && obj.Order != current.Order {
			_, err := db.Redis.Set(Key, obj.Order, 0).Result()
			return obj, current, err
		}
	}
	return obj, current, errors.New("not found")
}
