package repository

import (
	generalRepo "github.com/peteraba/d5/finder/repository/general"
	germanRepo "github.com/peteraba/d5/finder/repository/german"
	"gopkg.in/mgo.v2"
)

type QueryRepo interface {
	CreateDictionary(mgoSession *mgo.Session, dbName, collectionName string, query map[string]string) (interface{}, error)
}

func CreateRepo(isGerman bool) QueryRepo {
	if isGerman {
		return germanRepo.Repo{}
	}

	return generalRepo.Repo{}
}
