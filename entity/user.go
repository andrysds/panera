package entity

import (
	"github.com/andrysds/panera/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User represent users document
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Birthday string
	Role     string
}

// UserCollection represent user collection
type UserCollection struct {
	mgo.Collection
}

// Users represent users collection
func Users() *UserCollection {
	return &UserCollection{
		*db.DB.C("users"),
	}
}

// All returns all documents
func (u *UserCollection) All() (result []User, err error) {
	err = u.Find(nil).All(&result)
	return result, err
}

// FindOne returns a document by id
func (u *UserCollection) FindOne(id string) (result User, err error) {
	err = u.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

// InsertOne inserts a new document
func (u *UserCollection) InsertOne(new User) error {
	return u.Insert(new)
}

// UpdateOne upserts a document by id
func (u *UserCollection) UpdateOne(id string, update User) error {
	return u.UpdateId(bson.ObjectIdHex(id), update)
}

// RemoveOne removes a document by id
func (u *UserCollection) RemoveOne(id string) error {
	return u.RemoveId(bson.ObjectIdHex(id))
}
