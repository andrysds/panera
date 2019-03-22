package entity

import (
	"github.com/andrysds/panera/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Users represent users collection
func Users() *mgo.Collection {
	return db.DB.C("users")
}

// User represent users document
type User struct {
	_Id      bson.ObjectId
	Name     string
	Username string
	Birthday string
	Role     string
}
