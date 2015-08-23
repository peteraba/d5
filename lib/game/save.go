package game

import "gopkg.in/mgo.v2"

func SaveAnswer(game Game, mgoDb *mgo.Database, collectionName string) error {
	var (
		err        error
		collection *mgo.Collection
	)

	collection = mgoDb.C(collectionName)

	err = collection.Insert(game)

	return err
}
