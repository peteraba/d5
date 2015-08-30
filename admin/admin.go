package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
	"github.com/peteraba/d5/lib/admin"
	"github.com/peteraba/d5/lib/mongo"
	"gopkg.in/mgo.v2"
)

const (
	d5DbHostEnv = "D5_DBHOST"
	d5DbNameEnv = "D5_DBNAME"
)

func parseEnvs() (string, string) {
	dbHost := os.Getenv(d5DbHostEnv)

	dbName := os.Getenv(d5DbNameEnv)

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

	router.POST("/user", admin.CreateUser)
	router.PATCH("/user", admin.UpdateUser)
	router.DELETE("/user", admin.DeleteUser)
	router.GET("/user", admin.ReadUser)

	router.POST("/game", admin.CreateGame)
	router.PATCH("/game", admin.UpdateGame)
	router.DELETE("/game", admin.DeleteGame)
	router.GET("/game", admin.ReadGame)

	router.POST("/game-user", admin.CreateUpdateUserGame)
	router.PATCH("/game-user", admin.CreateUpdateUserGame)
	router.DELETE("/game-user", admin.DeleteUserGame)

	router.Run(fmt.Sprintf(":%d", port))
}
