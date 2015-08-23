package game

import "gopkg.in/mgo.v2"

func SaveAnswer(game Game, hostName, dbName, collectionName string) error {
	var (
		err        error
		collection *mgo.Collection
	)

	session, err := mgo.Dial(hostName)
	if err != nil {
		return err
	}
	defer session.Close()

	collection = session.DB(dbName).C(collectionName)

	err = collection.Insert(game)

	return err
}
