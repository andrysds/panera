package entity

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Standups represents standup collection
var Standups *Collection

// Standup represents standup document
type Standup struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"user_id"`
	State     string
	Timestamp time.Time
}

// User gets user object of the standup object
func (s *Standup) User() User {
	var result User
	Users.FindOne(s.UserID.Hex(), &result)
	return result
}

// SetDone sets state info to "done"
func (s *Standup) SetDone() error {
	s.State = "done"
	return Standups.UpdateOne(s.ID.Hex(), s)
}
