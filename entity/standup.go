package entity

import "github.com/globalsign/mgo/bson"

// Standups represents standup collection
var Standups *Collection

// Standup represents standup document
type Standup struct {
	ID     bson.ObjectId `bson:"_id"`
	UserID bson.ObjectId
	State  string
}
