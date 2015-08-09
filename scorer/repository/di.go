package repository

import (
	generalRepo "github.com/peteraba/d5/scorer/repository/general"
	germanRepo "github.com/peteraba/d5/scorer/repository/german"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QueryRepo interface {
	SetDb(db *mgo.Database)
	UpdateWord(collectionName string, objectId bson.ObjectId, result int) (bool, error)
}

func CreateRepo(mgoSession *mgo.Session, dbName string, isGerman bool) QueryRepo {
	var (
		repo QueryRepo
	)

	if isGerman {
		repo = &germanRepo.Repo{}
	} else {
		repo = &generalRepo.Repo{}
	}

	mgoSession = mgoSession.Clone()

	repo.SetDb(mgoSession.DB(dbName))

	return repo
}
