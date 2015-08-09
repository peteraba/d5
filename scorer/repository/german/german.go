package german

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	general "github.com/peteraba/d5/lib/general"
	german "github.com/peteraba/d5/lib/german"
)

type Repo struct {
	Db *mgo.Database
}

func (r *Repo) UpdateWord(collectionName string, objectId bson.ObjectId, result int) (bool, error) {
	var (
		collection *mgo.Collection
		err        error
		superwords = []german.Superword{}
	)

	collection = r.Db.C(collectionName)

	err = collection.FindId(objectId).All(&superwords)
	if err != nil {
		log.Printf("Error while finding word with id: %v, err: %v\n", objectId, err)

		return false, err
	}

	if len(superwords) < 1 {
		log.Printf("Failed to find word with id: %v\n", objectId)

		return false, nil
	}

	//log.Printf("Adding score to word with id: %v\n", objectId)

	err = addScore(&superwords[0], result)
	if err != nil {
		log.Printf("Error while adding score: %v\n", err)

		return false, nil
	}

	err = collection.UpdateId(objectId, superwords[0])
	if err != nil {
		log.Printf("Error happend while updating word id: %v. Err: %v\n", objectId, err)

		return false, err
	}

	//log.Printf("No error happend while updating word id: %v, score: %d.\n", objectId, result)

	return true, nil
}

func addScore(superword *german.Superword, result int) error {
	var (
		err      error
		newScore *general.Score
	)

	newScore, err = general.NewScore(result)
	if err != nil {
		return err
	}

	superword.DefaultWord.Scores = append(superword.DefaultWord.Scores, newScore)

	return nil
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
