package repository

import (
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QueryRepo interface {
	SetDb(db *mgo.Database)
	FetchWord(collectionName string, objectId bson.ObjectId) (entity.Word, error)
	FetchDictionary(collectionName string, query map[string]string) (interface{}, error)
	FilterDictionary(limit int) (interface{}, error)
	UpdateWord(collectionName string, objectId bson.ObjectId, data interface{}) error
}

func CreateRepo(mgoSession *mgo.Session, dbName string, isGerman bool) QueryRepo {
	var (
		repo QueryRepo
	)

	repo = &Repo{}

	mgoSession = mgoSession.Clone()

	repo.SetDb(mgoSession.DB(dbName))

	return repo
}
