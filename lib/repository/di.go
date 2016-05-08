package repository

import (
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QueryRepo interface {
	SetDb(db *mgo.Database)
	FetchWord(collectionName string, objectId bson.ObjectId) (entity.Word, error)
	FetchDictionary(collectionName string, query bson.M) (interface{}, error)
	FilterDictionary(limit int) (interface{}, error)
	UpdateWord(collectionName string, objectId bson.ObjectId, data interface{}) error
}

func CreateRepo(mgoDb *mgo.Database) QueryRepo {
	var (
		repo QueryRepo
	)

	repo = &Repo{}

	repo.SetDb(mgoDb)

	return repo
}
