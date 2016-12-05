package util

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func MgoDb(mgoDb *mgo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mgoDb", mgoDb)

		c.Next()
	}
}
