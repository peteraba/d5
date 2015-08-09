package util

import "gopkg.in/mgo.v2/bson"

func HexToObjectId(wordId string) *bson.ObjectId {
	var (
		objectId bson.ObjectId
	)

	if !bson.IsObjectIdHex(wordId) {
		return nil
	}

	objectId = bson.ObjectIdHex(wordId)

	return &objectId
}
