package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
	"github.com/peteraba/d5/lib/admin"
	"github.com/peteraba/d5/lib/mongo"
)

const (
	dbhost_env = "D5_DBHOST"
	dbname_env = "D5_DBNAME"
)

func parseEnvs() (string, string) {
	dbHost := os.Getenv(dbhost_env)

	dbName := os.Getenv(dbname_env)

	return dbHost, dbName
}

func parseFlags() (int, bool) {
	port := flag.Int("port", 17174, "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *port, *debug
}

func MgoDb(mgoDb *mgo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mgoDb", mgoDb)

		c.Next()
	}
}

func main() {
	port, debug := parseFlags()

	dbHost, dbName := parseEnvs()
	mgoDb, err := mongo.CreateMgoDb(dbHost, dbName)
	if err != nil {
		log.Fatalln(err)
	}

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creates a gin router with default middlewares:
	// logger and recovery (crash-free) middlewares
	router := gin.Default()

	router.Use(MgoDb(mgoDb))

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
