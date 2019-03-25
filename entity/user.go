package entity

import "github.com/globalsign/mgo/bson"

// Users represent user collection
var Users *Collection

// User represent user document
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Birthday string
	Role     string
}
