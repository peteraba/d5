package mongo

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2"
)

func TestCreateMongoDb(t *testing.T) {
	var (
		mgoSession = mgo.Session{}
		mgoDb      = mgo.Database{}
	)

	SetMgoSession(&mgoSession)
	SetMgoDb(&mgoDb)

	db, err := CreateMgoDb("foo", "bar")

	if err != nil {
		t.Fatalf("Error happened while retrieving database: %v.", err)
	}

	if !reflect.DeepEqual(mgoDb, *db) {
		t.Fatalf("Returned database is not the mock database.\nMock: %v\nGot: %v\n", mgoDb, *db)
	}

	t.Log(1, "test cases")
}
