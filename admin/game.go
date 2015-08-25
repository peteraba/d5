package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameForm struct {
	name     string `json:"username" binding:"required"`
	route    string `json:"route"`
	url      string `json:"url"`
	isSystem bool   `json:"is_system"`
}

func createGame(c *gin.Context) {
	var gameForm GameForm

	c.Bind(&gameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func deleteGame(c *gin.Context) {
	var gameForm GameForm

	c.Bind(&gameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func updateGame(c *gin.Context) {
	var gameForm GameForm

	c.Bind(&gameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func readGame(c *gin.Context) {
	var gameForm GameForm

	c.Bind(&gameForm)
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}
