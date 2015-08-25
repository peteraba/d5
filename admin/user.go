package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	username string `json:"username" binding:"required"`
	maxWords int    `json:"max-words"`
}

func createUser(c *gin.Context) {
	var userForm UserForm

	c.Bind(&userForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func deleteUser(c *gin.Context) {
	var userForm UserForm

	c.Bind(&userForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func updateUser(c *gin.Context) {
	var userForm UserForm

	c.Bind(&userForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func readUser(c *gin.Context) {
	var userForm UserForm

	c.Bind(&userForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func listGameUsers(c *gin.Context) {
	var userForm UserForm

	c.Bind(&userForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}
