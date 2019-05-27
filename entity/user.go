package entity

import (
	"github.com/globalsign/mgo/bson"
)

// Users represent user collection
var Users *Collection

// User represent user document
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Birthday string
	Active   bool
}

// AllUsers returns all user records
func AllUsers() (users []*User, err error) {
	err = Users.All("name", &users)
	return users, err
}

// AddUserToStandups adds user to standup records
func AddUserToStandups(id string) error {
	newStandup := Standup{
		ID:     bson.NewObjectId(),
		UserID: bson.ObjectIdHex(id),
		State:  "undone",
	}
	return Standups.InsertOne(newStandup)
}
