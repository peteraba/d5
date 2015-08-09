package mongo

import (
	"os"

	"gopkg.in/mgo.v2"
)

const (
	d5_dbhost_env = "D5_DBHOST"
	d5_dbname_env = "D5_DBNAME"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

func ParseEnvs() (string, string) {
	// Mongo database host
	hostname := os.Getenv(d5_dbhost_env)

	// Mongo database name
	dbName := os.Getenv(d5_dbname_env)

	return hostname, dbName
}

func SetMgoSession(mgoSession *mgo.Session) {
	session = mgoSession
}

func GetMgoSession(url string) (*mgo.Session, error) {
	var (
		err error
	)

	if session != nil {
		return session, nil
	}

	session, err = mgo.Dial(url)

	return session, err
}

func SetMgoDb(mgoDatabase *mgo.Database) {
	db = mgoDatabase
}

func GetMgoDb(mgoSession *mgo.Session, dbName string) *mgo.Database {
	if db != nil {
		return db
	}

	mgoSession = mgoSession.Clone()

	db = mgoSession.DB(dbName)

	return db
}

func CreateMgoDb() (*mgo.Database, error) {
	hostname, dbName := ParseEnvs()

	session, err := GetMgoSession(hostname)
	if err != nil {
		return nil, err
	}

	return GetMgoDb(session, dbName), nil
}
