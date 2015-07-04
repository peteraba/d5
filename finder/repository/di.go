package repository

import (
	generalRepo "github.com/peteraba/d5/finder/repository/general"
	germanRepo "github.com/peteraba/d5/finder/repository/german"
	"gopkg.in/mgo.v2"
)

type QueryRepo interface {
	SetDb(db *mgo.Database)
	CreateDictionary(collectionName string, query map[string]string) (interface{}, error)
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
