package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	gameColl = "game"
)

type Game struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Route    string `form:"route" json:"route" binding:"required"`
	Url      string `form:"url" json:"url" binding:"required"`
	IsSystem bool   `form:"is-system" json:"isSystem" bson:"isSystem,omitempty"`
}

func CreateGame(c *gin.Context) {
	var (
		err  error
		game Game
	)

	mgoDb := c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection := mgoDb.C(gameColl)

	if c.Bind(&game) != nil {
		BadRequest(c)

		return
	}

	err = mgoCollection.Insert(game)
	if err != nil {
		InternalServerError(c, err, "Inserting game failed")

		return
	}

	Ok(c)
}

func UpdateGame(c *gin.Context) {
	var (
		err           error
		game          Game
		id            string
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		objectId      *bson.ObjectId
	)

	id = c.PostForm("id")
	if id == "" {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(id)

	mgoDb = c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection = mgoDb.C(gameColl)

	if c.Bind(&game) != nil {
		BadRequest(c)

		return
	}

	err = mgoCollection.UpdateId(objectId, game)
	if err != nil {
		InternalServerError(c, err, "Updating game failed")

		return
	}

	Ok(c)
}

func DeleteGame(c *gin.Context) {
	var (
		err           error
		id            string
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		objectId      *bson.ObjectId
	)

	id = c.PostForm("id")
	if id == "" {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(id)

	mgoDb = c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection = mgoDb.C(gameColl)

	err = mgoCollection.RemoveId(objectId)
	if err != nil {
		InternalServerError(c, err, "Removing game failed")

		return
	}

	Ok(c)
}

func ReadGame(c *gin.Context) {
	var (
		err    error
		result []Game
	)

	mgoDb := c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection := mgoDb.C(gameColl)

	err = mgoCollection.Find(bson.M{}).All(&result)
	if err != nil {
		InternalServerError(c, err, "Listing games failed")

		return
	}

	OkWithData(c, result)
}
