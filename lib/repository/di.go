package repository

import (
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QueryRepo interface {
	FetchWord(collectionName string, objectId bson.ObjectId) (entity.Word, error)
	UpdateWord(collectionName string, objectId bson.ObjectId, data interface{}) error
	SetDb(db *mgo.Database)
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
