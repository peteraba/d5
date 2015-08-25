package admin

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func parseFlags() (int, bool) {
	port := flag.Int("port", 17172, "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *port, *debug
}

func main() {
	port, _ := parseFlags()

	// Creates a gin router with default middlewares:
	// logger and recovery (crash-free) middlewares
	router := gin.Default()

	router.POST("/user", createUser)
	router.DELETE("/user", deleteUser)
	router.PATCH("/user/", updateUser)
	router.GET("/user", readUser)

	router.POST("/user/game", createUserGame)
	router.DELETE("/user/game", deleteUserGame)
	router.PATCH("/user/game", updateUserGame)
	router.GET("/user/game", readUserGame)
	router.GET("/usergame", listGameUsers)

	router.POST("/game", createGame)
	router.DELETE("/game", deleteGame)
	router.PATCH("/game", updateGame)
	router.GET("/game", readGame)

	router.Run(fmt.Sprintf(":%d", port))
}
