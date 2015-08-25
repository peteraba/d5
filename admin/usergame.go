package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserGameForm struct {
	name     string `json:"username" binding:"required"`
	route    string `json:"route"`
	url      string `json:"url"`
	isSystem bool   `json:"is_system"`
}

func createUserGame(c *gin.Context) {
	var userGameForm UserGameForm

	c.Bind(&userGameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func deleteUserGame(c *gin.Context) {
	var userGameForm UserGameForm

	c.Bind(&userGameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func updateUserGame(c *gin.Context) {
	var userGameForm UserGameForm

	c.Bind(&userGameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func readUserGame(c *gin.Context) {
	var userGameForm UserGameForm

	c.Bind(&userGameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}
