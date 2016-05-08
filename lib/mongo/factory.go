package mongo

import "gopkg.in/mgo.v2"

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

func CreateMgoDb(hostname, dbName string) (*mgo.Database, error) {
	session, err := getMgoSession(hostname)
	if err != nil {
		return nil, err
	}

	return getMgoDb(session, dbName), nil
}

func getMgoSession(hostName string) (*mgo.Session, error) {
	var (
		err error
	)

	if session != nil {
		return session, nil
	}

	session, err = mgo.Dial(hostName)

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
