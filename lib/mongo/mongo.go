package mongo

import (
	"os"

	"gopkg.in/mgo.v2"
)

const (
	d5_dbhost_env = "D5_DBHOST"
	d5_dbname_env = "D5_DBNAME"
)

func ParseEnvs() (string, string) {
	// Mongo database host
	hostname := os.Getenv(d5_dbhost_env)

	// Mongo database name
	dbName := os.Getenv(d5_dbname_env)

	return hostname, dbName
}

func GetMgoSession(url string) (*mgo.Session, error) {
	var (
		err     error
		session *mgo.Session
	)

	session, err = mgo.Dial(url)
	if err != nil {
		return session, err
	}

	//session.SetMode(mgo.Monotonic, true)
	//session.SetSafe(&mgo.Safe{})

	return session, err
}

func GetMgoDb(mgoSession *mgo.Session, dbName string) *mgo.Database {
	mgoSession = mgoSession.Clone()

	return mgoSession.DB(dbName)
}

func CreateMgoDb() (*mgo.Database, error) {
	hostname, dbName := ParseEnvs()

	session, err := GetMgoSession(hostname)
	if err != nil {
		return nil, err
	}

	return GetMgoDb(session, dbName), nil
}
