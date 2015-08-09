package general

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repo struct {
	Db *mgo.Database
}

func (r *Repo) UpdateWord(collectionName string, object bson.ObjectId, result int) (bool, error) {
	return false, nil
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
