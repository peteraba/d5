package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func parseFlags() (int, bool, string, string) {
	port := flag.Int("port", 17182, "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	finder := flag.String("finder", "http://localhost:17171/", "Finder address")

	scorer := flag.String("scorer", "http://localhost:17172/", "Scorer address")

	flag.Parse()

	return *port, *debug, *finder, *scorer
}

func main() {
	port, debug, finder, scorer := parseFlags()

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/game", makeGameHandle(finder))
	router.POST("/answer", makeCheckAnswerHandle(scorer))

	router.Run(fmt.Sprintf(":%d", port))
}

func makeGameHandle(finder string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(200, finder)
	}
}

func makeCheckAnswerHandle(scorer string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(200, scorer)
	}
}
