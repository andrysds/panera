package db

import (
	"log"
	"os"

	"github.com/andrysds/clarity"
	"github.com/globalsign/mgo"
)

// DB represents application database
var DB *mgo.Database

// Init initializes application database connection
func Init() {
	mongoURL := os.Getenv("MONGODB_URI")
	session, err := mgo.Dial(mongoURL)
	clarity.PanicIfError(err, "error on connecting to database")
	DB = session.DB("")
	log.Println("* DB initialized")
}
