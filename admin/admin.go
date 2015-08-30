package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/admin"
	"github.com/peteraba/d5/lib/mongo"
)

const (
	d5_dbhost_env = "D5_DBHOST"
	d5_dbname_env = "D5_DBNAME"
)

func parseEnvs() (string, string) {
	dbHost := os.Getenv(d5_dbhost_env)

	dbName := os.Getenv(d5_dbname_env)

	return dbHost, dbName
}

func parseFlags() (int, bool) {
	port := flag.Int("port", 17174, "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *port, *debug
}

func MgoDb(dbHost, dbName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mgoDb, err := mongo.CreateMgoDb(dbHost, dbName)
		if err != nil {
			log.Print(err)
		}

		c.Set("mgoDb", mgoDb)

		c.Next()
	}
}

func main() {
	port, _ := parseFlags()

	dbHost, dbName := parseEnvs()

	// Creates a gin router with default middlewares:
	// logger and recovery (crash-free) middlewares
	router := gin.Default()

	router.Use(MgoDb(dbHost, dbName))

	router.POST("/user", admin.CreateUser)
	router.DELETE("/user", admin.DeleteUser)
	router.PATCH("/user", admin.UpdateUser)
	router.GET("/user", admin.ReadUser)

	router.POST("/game", admin.CreateGame)
	router.DELETE("/game", admin.DeleteGame)
	router.PATCH("/game", admin.UpdateGame)
	router.GET("/game", admin.ReadGame)

	router.Run(fmt.Sprintf(":%d", port))
}
