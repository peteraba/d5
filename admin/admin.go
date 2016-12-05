package main

import (
	"fmt"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	admin "github.com/peteraba/d5/admin/lib"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/util"
)

const name = "admin"
const version = "0.1"
const usage = `
Admin supports CLI and Server mode.

In CLI mode it expects input data on standard input, in server mode as raw POST body

Usage:
  admin [--server] [--port=<n>] [--debug] [--user=<s>]
  admin -h | --help
  admin -v | --version

Options:
  -s, --server    run in server mode
  -p, --port=<n>  port to open (server mode only) [default: 10310]
  -d, --debug     skip ticks and generate fake data concurrently
  -v, --version   show version information
  -h, --help      show help information
  -u, --user=<s>  user the data belongs to (cli mode only)
`

/**
 * MAIN
 */

func main() {
	cliArguments := util.GetCliArguments(usage, name, version)
	_, port, isDebug := util.GetServerOptions(cliArguments)

	mgoDb, err := mongo.CreateMgoDbFromEnvs()
	util.LogFatalfMsg(err, "MongoDB database could not be created: %v", true)

	startServer(port, mgoDb, isDebug)
}

/**
 * SERVER
 */

func startServer(port int, mgoDb *mgo.Database, isDebug bool) {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creates a gin router with default middlewares:
	// logger and recovery (crash-free) middlewares
	router := gin.Default()

	router.Use(util.MgoDb(mgoDb))

	router.GET("/game", admin.ReadGame)
	router.POST("/game", admin.CreateGame)
	router.PATCH("/game/:id", admin.UpdateGame)
	router.DELETE("/game/:id", admin.DeleteGame)

	router.GET("/user/:id", admin.ReadUser)
	router.POST("/user", admin.CreateUser)
	router.PATCH("/user/:id", admin.UpdateUser)
	router.DELETE("/user/:id", admin.DeleteUser)

	router.POST("/game-user", admin.CreateUpdateUserGame)
	router.PATCH("/game-user/:userId/:gameName", admin.CreateUpdateUserGame)
	router.DELETE("/game-user/:userId/:gameName", admin.DeleteUserGame)

	router.Run(fmt.Sprintf(":%d", port))
}
