package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

func SetResultIndexes(mgoCollection *mgo.Collection) error {
	index := mgo.Index{
		Key:         []string{"createdAt"},
		Unique:      false,
		DropDups:    false,
		Background:  true,
		Sparse:      true,
		ExpireAfter: time.Hour,
		Name:        "expire",
	}

	return mgoCollection.EnsureIndex(index)
}
