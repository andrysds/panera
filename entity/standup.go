package entity

import (
	"log"
	"time"

	"github.com/andrysds/panera/connection"
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

// GetTodayStandup pick one random standup record
func GetTodayStandup() (result Standup) {
	id, err := connection.Redis.Get("panera:standup").Result()
	log.Println(id, err)
	if len(id) == 0 {
		Standups.Pipe([]bson.M{{"$sample": bson.M{"size": 1}}}).One(&result)
		err := connection.Redis.Set("panera:standup", result.ID.Hex(), 0).Err()
		log.Println(err)
	} else {
		Standups.FindOne(id, &result)
	}
	return result
}

// GetStandupList returns all standup records
func GetStandupList() (results []Standup) {
	err := Standups.Find(bson.M{}).All(&results)
	log.Println(err)
	return results
}
