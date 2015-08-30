package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserGame struct {
	Game    string      `json:"game"`
	Weight  int         `json:"weight"`
	Options interface{} `json:"options"`
}

type UserGameForm struct {
	UserId  string      `form:"user-id" binding:"required"`
	Game    string      `form:"game" binding:"required"`
	Weight  int         `form:"weight"`
	Options interface{} `form:"options"`
}

func CreateUpdateUserGame(c *gin.Context) {
	var (
		err           error
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		userGameForm  UserGameForm
		objectId      *bson.ObjectId
		result        []User
		user          User
	)

	if c.Bind(&userGameForm) != nil {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(userGameForm.UserId)

	mgoDb = c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection = mgoDb.C(userColl)

	err = mgoCollection.FindId(objectId).All(&result)
	if err != nil {
		InternalServerError(c, err, "Finding user failed")

		return
	}

	if len(result) < 1 {
		Accepted(c)

		return
	}

	user.Games[userGameForm.Game] = UserGame{userGameForm.Game, userGameForm.Weight, userGameForm.Options}

	err = mgoCollection.UpdateId(objectId, user)
	if err != nil {
		InternalServerError(c, err, "Updating user failed")

		return
	}

	Ok(c)
}

func DeleteUserGame(c *gin.Context) {
	var (
		err           error
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		userGameForm  UserGameForm
		objectId      *bson.ObjectId
		result        []User
		user          User
	)

	if c.Bind(&userGameForm) != nil {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(userGameForm.UserId)

	mgoDb = c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection = mgoDb.C(userColl)

	err = mgoCollection.FindId(objectId).All(&result)
	if err != nil {
		InternalServerError(c, err, "Finding user failed")

		return
	}

	if len(result) < 1 {
		Accepted(c)

		return
	}

	if _, ok := user.Games[userGameForm.Game]; !ok {
		Accepted(c)

		return
	}

	delete(user.Games, userGameForm.Game)

	err = mgoCollection.UpdateId(objectId, user)
	if err != nil {
		InternalServerError(c, err, "Updating user failed")

		return
	}

	Ok(c)
}
