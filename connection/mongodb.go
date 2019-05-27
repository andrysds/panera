package connection

import (
	"log"
	"os"

	"github.com/andrysds/clarity"
	"github.com/globalsign/mgo"
)

// MongoDB represents MongoDB connection
var MongoDB *mgo.Database

// InitMongoDB initializes MongoDB connection
func InitMongoDB() {
	mongoURL := os.Getenv("MONGODB_URI")
	session, err := mgo.Dial(mongoURL)
	clarity.PanicIfError(err, "error on connecting to database")
	MongoDB = session.DB("")
	log.Println("* MongoDB initialized")
}

// MongoNotFoundErr is message for not found error from mongo
const MongoNotFoundErr = "not found"
