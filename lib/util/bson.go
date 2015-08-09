package util

import (
	"encoding/hex"

	"gopkg.in/mgo.v2/bson"
)

func HexToObjectId(wordId string) *bson.ObjectId {
	var (
		wordIdBytes []byte
		err         error
		objectId    bson.ObjectId
	)

	wordIdBytes, err = hex.DecodeString(wordId)
	if err != nil {
		return nil
	}

	objectId = bson.ObjectId(wordIdBytes)

	return &objectId
}
