package entity

import (
	"github.com/andrysds/panera/connection"
	"github.com/globalsign/mgo/bson"
)

const (
	// StandupKey is redis key for today standup cache
	StandupKey = "panera:standup"

	// StandupStateDone is state for done standup
	StandupStateDone = "done"
	// StandupStateSkipped is state for skipped standup
	StandupStateSkipped = "skipped"
	// StandupStateUndone is state for undone standup
	StandupStateUndone = "undone"
)

// Standups represents standup collection
var Standups *Collection

// Standup represents standup document
type Standup struct {
	ID     bson.ObjectId `bson:"_id"`
	UserID bson.ObjectId `bson:"user_id"`
	State  string
}

// User gets user object of the standup object
func (s *Standup) User() (result *User, err error) {
	err = Users.FindOne(s.UserID.Hex(), &result)
	return result, err
}

// SetState sets state info
func (s *Standup) SetState(state string) error {
	s.State = state
	return Standups.UpdateOne(s.ID.Hex(), s)
}

// GetStandup returns current standup record
func GetStandup() (*Standup, error) {
	var result *Standup
	id, err := connection.Redis.Get(StandupKey).Result()
	if err != nil {
		err = GetNewStandup(&result)
		if err != nil && err.Error() == connection.MongoNotFoundErr {
			if err = NewPeriodStandup(); err == nil {
				err = GetNewStandup(&result)
			}
		}
		if err == nil {
			err = connection.Redis.Set(StandupKey, result.ID.Hex(), 0).Err()
		}
	} else {
		err = Standups.FindOne(id, &result)
	}
	return result, err
}

// GetNewStandup pick random standup record from database
func GetNewStandup(container interface{}) error {
	return Standups.Pipe([]bson.M{
		{"$match": bson.M{"state": "undone"}},
		{"$sample": bson.M{"size": 1}},
	}).One(container)
}

// GetStandupList returns all standup records
func GetStandupList() ([]*Standup, error) {
	var results []*Standup
	err := Standups.All("state", &results)
	return results, err
}

// NewDayStandup sets today standup state done and clear the cache(redis)
func NewDayStandup() error {
	standup, err := GetStandup()
	if err == nil {
		if err = standup.SetState(StandupStateDone); err == nil {
			if _, err = connection.Redis.Del(StandupKey).Result(); err == nil {
				var standups []*Standup
				if err = Standups.Find(bson.M{"state": StandupStateSkipped}).All(&standups); err == nil {
					for _, s := range standups {
						s.SetState(StandupStateUndone)
					}
				}
			}
		}
	}
	return err
}

// NewPeriodStandup reset standup data and setup new period standup
func NewPeriodStandup() error {
	_, err := Standups.RemoveAll(bson.M{"state": StandupStateDone})
	if err == nil {
		var users []*User
		if err = Users.Find(bson.M{"active": true}).All(&users); err == nil {
			for _, u := range users {
				if err := AddUserToStandups(u.ID.Hex()); err != nil {
					return err
				}
			}
		}
	}
	return err
}

// SkipStandup sets today standup state skipped and return new standup
func SkipStandup() (*Standup, *Standup, error) {
	var standup *Standup
	skipped, err := GetStandup()
	if err == nil {
		err = skipped.SetState(StandupStateSkipped)
		if err == nil {
			_, err = connection.Redis.Del(StandupKey).Result()
			if err == nil {
				standup, err = GetStandup()
			}
		}
	}
	return standup, skipped, err
}
