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

	for i := current.Order + 1; i < len(objs); i++ {
		if objs[i].State != "1" {
			_, err := db.Redis.Set(Key, i, 0).Result()
			return objs[i], current, err
		}
	}
	return obj, current, errors.New("not found")
}
