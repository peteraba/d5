package mongo

import (
	"os"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2"
)

func TestParseEnvs(t *testing.T) {
	os.Setenv("D5_DBHOST", "foo")
	os.Setenv("D5_DBNAME", "bar")

	hostname, dbName := ParseEnvs()

	if hostname != "foo" {
		t.Fatalf("Hostname is not 'foo', it's %s.", hostname)
	}

	if dbName != "bar" {
		t.Fatalf("Database name is not 'bar', it's %s.", dbName)
	}

	t.Log(1, "test cases")
}

func TestCreateMongoDb(t *testing.T) {
	var (
		mgoSession = mgo.Session{}
		mgoDb      = mgo.Database{}
	)

	os.Setenv("D5_DBHOST", "foo")
	os.Setenv("D5_DBNAME", "bar")

	SetMgoSession(&mgoSession)
	SetMgoDb(&mgoDb)

	db, err := CreateMgoDb()

	if err != nil {
		t.Fatalf("Error happened while retrieving database: %v.", err)
	}

	if !reflect.DeepEqual(mgoDb, *db) {
		t.Fatalf("Returned database is not the mock database.\nMock: %v\nGot: %v\n", mgoDb, *db)
	}

	t.Log(1, "test cases")
}
