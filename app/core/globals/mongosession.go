package globals

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/next/app/core/cfg"
	"github.com/sknv/next/app/lib/mongo"
)

var (
	mongoSession *mgo.Session
)

func InitMongoSession(config *cfg.Config) {
	dialInfo := &mgo.DialInfo{
		Addrs:    config.MongoAddrs,
		Database: config.MongoDatabase,
		Username: config.MongoUsername,
		Password: config.MongoPassword,
		Timeout:  config.MongoTimeout,
		Source:   "admin", // authenticate against "admin" database
	}
	mongoSession = mongo.MustDial(dialInfo)
}

// GetMongoSession returns a global mgo.Session.
func GetMongoSession() *mgo.Session {
	return mongoSession
}
