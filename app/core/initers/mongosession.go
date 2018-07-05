package initers

import (
	"github.com/globalsign/mgo"

	"github.com/sknv/next/app/lib/mongo"
)

var (
	mongoSession *mgo.Session
)

func init() {
	cfg := GetConfig()
	dialInfo := &mgo.DialInfo{
		Addrs:    cfg.MongoAddrs,
		Database: cfg.MongoDatabase,
		Username: cfg.MongoUsername,
		Password: cfg.MongoPassword,
		Timeout:  cfg.MongoTimeout,
		Source:   "admin", // authenticate against "admin" database
	}
	mongoSession = mongo.MustDial(dialInfo)
}

// GetMongoSession returns a global mgo.Session.
func GetMongoSession() *mgo.Session {
	return mongoSession
}
