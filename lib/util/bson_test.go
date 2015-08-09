package util

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestHexToObjectIdConvertsHexadecimalStringToObjectId(t *testing.T) {
	var (
		hexa     = "55c7177c288a2139ea45ebe4"
		objectId *bson.ObjectId
	)

	objectId = HexToObjectId(hexa)

	if hexa != objectId.Hex() {
		t.Fatalf("Returned ObjectId has a different id from the one used to create it. Expected: %s, found: %s", hexa, objectId.Hex())
	}

	t.Log(1, "test cases")
}

func TestHexToObjectIdReturnsNilWhenStringIsInvalidObjectId(t *testing.T) {
	var (
		hexa     = "p5c7177c288a2139ea45ebe4"
		objectId *bson.ObjectId
	)

	objectId = HexToObjectId(hexa)

	if objectId != nil {
		t.Fatalf("Hello")
	}

	t.Log(1, "test cases")
}
