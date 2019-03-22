package db

import (
	"log"
	"os"

	"github.com/andrysds/clarity"
	"github.com/globalsign/mgo"
)

// DB represents project database
var DB *mgo.Database

// Init initiates databse connection
func Init() {
	mongoURL := os.Getenv("MONGODB_URI")
	log.Println(mongoURL)
	session, err := mgo.Dial(mongoURL)
	clarity.PanicIfError(err, "error on connecting to database")
	DB = session.DB("")
	log.Println("* DB initialized")
}
