package repository

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
)

type Repo struct {
	Db *mgo.Database
}

func (r *Repo) FetchWord(collectionName string, objectId bson.ObjectId) (entity.Word, error) {
	var (
		err        error
		collection *mgo.Collection
		superwords = []german.Superword{}
	)

	collection = r.Db.C(collectionName)

	err = collection.FindId(objectId).All(&superwords)
	if err != nil {
		log.Printf("Error while finding word with id: %v, err: %v\n", objectId, err)

		return nil, err
	}

	if len(superwords) < 1 {
		log.Printf("Failed to find word with id: %v\n", objectId)

		return nil, nil
	}

	words := german.SuperwordsToWords(superwords)

	return words[0], nil
}

func (r *Repo) UpdateWord(collectionName string, objectId bson.ObjectId, data interface{}) error {
	var (
		err        error
		collection *mgo.Collection
	)

	collection = r.Db.C(collectionName)

	err = collection.UpdateId(objectId, data)
	if err != nil {
		log.Printf("Error happend while updating word id: %v. Err: %v\n", objectId, err)

		return err
	}

	return nil
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
