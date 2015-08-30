package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	userColl = "user"
)

type User struct {
	Username string   `form:"username" json:"username" binding:"required"`
	MaxWords int      `form:"max-words" json:"maxWords"`
	Games    []string `form:"games" json:"games" bson:"games,omitempty"`
}

func CreateUser(c *gin.Context) {
	var (
		err  error
		user User
	)

	mgoDb := c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection := mgoDb.C(userColl)

	if c.Bind(&user) != nil {
		BadRequest(c)

		return
	}

	err = mgoCollection.Insert(user)
	if err != nil {
		InternalServerError(c, err, "Inserting user failed")

		return
	}

	Ok(c)
}

func UpdateUser(c *gin.Context) {
	var (
		err           error
		userNew       User
		id            string
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		result        []User
		userOld       User
		objectId      *bson.ObjectId
	)

	id = c.PostForm("id")
	if id == "" {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(id)

	mgoDb = c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection = mgoDb.C(userColl)

	if c.Bind(&userNew) != nil {
		BadRequest(c)

		return
	}

	err = mgoCollection.FindId(objectId).All(&result)
	if err != nil {
		InternalServerError(c, err, "Finding user failed")

		return
	}

	if len(result) < 1 {
		Accepted(c)

		return
	}

	userOld = result[0]
	if c.PostForm("username") != "" {
		userOld.Username = userNew.Username
	}
	if c.PostForm("max-words") != "" {
		userOld.MaxWords = userNew.MaxWords
	}
	if c.PostForm("games") != "" {
		userOld.Games = userNew.Games
	}

	err = mgoCollection.UpdateId(objectId, userOld)
	if err != nil {
		InternalServerError(c, err, "Updating user failed")

		return
	}

	Ok(c)
}

func DeleteUser(c *gin.Context) {
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

	mgoCollection = mgoDb.C(userColl)

	err = mgoCollection.RemoveId(objectId)
	if err != nil {
		InternalServerError(c, err, "Removing user failed")

		return
	}

	Ok(c)
}

func ReadUser(c *gin.Context) {
	var (
		err      error
		id       string
		result   []User
		objectId *bson.ObjectId
	)

	id = c.PostForm("id")
	if id == "" {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(id)

	mgoDb := c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection := mgoDb.C(userColl)

	err = mgoCollection.FindId(objectId).All(&result)
	if err != nil {
		InternalServerError(c, err, "Finding user failed")

		return
	}

	OkWithData(c, result)
}
