package entity

import (
	"log"

	"github.com/andrysds/panera/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Collection is abstract type of collection connection
type Collection struct{ mgo.Collection }

// All returns all documents
func (c *Collection) All(sort string, result interface{}) error {
	return c.Find(nil).Sort(sort).All(result)
}

// FindOne returns a document by id
func (c *Collection) FindOne(id string, result interface{}) error {
	return c.FindId(bson.ObjectIdHex(id)).One(result)
}

// InsertOne inserts a new document
func (c *Collection) InsertOne(new interface{}) error {
	return c.Insert(new)
}

// UpdateOne upserts a document by id
func (c *Collection) UpdateOne(id string, update interface{}) error {
	return c.UpdateId(bson.ObjectIdHex(id), update)
}

// RemoveOne removes a document by id
func (c *Collection) RemoveOne(id string) error {
	return c.RemoveId(bson.ObjectIdHex(id))
}

// InitCollection initiates collection connections
func InitCollection() {
	Standups = &Collection{*db.DB.C("standups")}
	Users = &Collection{*db.DB.C("users")}
	log.Println("* Collection initialized")
}
