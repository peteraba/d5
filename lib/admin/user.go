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
	Username string              `json:"username"`
	MaxWords int                 `json:"maxWords"`
	Games    map[string]UserGame `json:"userGames"`
}

type UserForm struct {
	Id       string `form:"id"`
	Username string `form:"username"`
	MaxWords int    `form:"max-words"`
}

func CreateUser(c *gin.Context) {
	var (
		err      error
		user     User
		userForm UserForm
	)

	mgoDb := c.MustGet("mgoDb").(*mgo.Database)

	mgoCollection := mgoDb.C(userColl)

	if c.Bind(&userForm) != nil {
		BadRequest(c)

		return
	}

	user = User{userForm.Username, userForm.MaxWords, make(map[string]UserGame)}

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
		userForm      UserForm
		mgoDb         *mgo.Database
		mgoCollection *mgo.Collection
		result        []User
		user          User
		objectId      *bson.ObjectId
	)

	if c.Bind(&userForm) != nil {
		BadRequest(c)

		return
	}

	objectId = util.HexToObjectId(userForm.Id)

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

	user = result[0]
	if c.PostForm("username") != "" {
		user.Username = userForm.Username
	}
	if c.PostForm("max-words") != "" {
		user.MaxWords = userForm.MaxWords
	}

	err = mgoCollection.UpdateId(objectId, user)
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
