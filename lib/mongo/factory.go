package mongo

import (
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

func SetMgoSession(mgoSession *mgo.Session) {
	session = mgoSession
}

func SetMgoDb(mgoDatabase *mgo.Database) {
	db = mgoDatabase
}

func CreateMgoDbFromEnvs() (*mgo.Database, error) {
	dbHost, dbName := ParseDbEnvs()
	util.LogMsg("Missing environment variable: D5_DBHOST", dbHost == "", util.IS_FATAL)
	util.LogMsg("Missing environment variable: D5_DBNAME", dbName == "", util.IS_FATAL)

	return CreateMgoDb(dbHost, dbName)
}

func CreateMgoDb(dbHost, dbName string) (*mgo.Database, error) {
	session, err := getMgoSession(dbHost)
	if err != nil {
		return nil, err
	}

	return getMgoDb(session, dbName), nil
}

func getMgoSession(dbHost string) (*mgo.Session, error) {
	var (
		err error
	)

	if session != nil {
		return session, nil
	}

	session, err = mgo.Dial(dbHost)

	return session, err
}

func getMgoDb(mgoSession *mgo.Session, dbName string) *mgo.Database {
	if db != nil {
		return db
	}

	mgoSession = mgoSession.Clone()

	db = mgoSession.DB(dbName)

	return db
}
