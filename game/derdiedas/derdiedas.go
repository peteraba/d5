package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/game"
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
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
	port, debug, finderUrl, scorerUrl := parseFlags()

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/game/:user", makeGameHandle(finderUrl))
	router.POST("/answer/:user", makeCheckAnswerHandle(finderUrl, scorerUrl))

	router.Run(fmt.Sprintf(":%d", port))
}

func makeGameHandle(finderUrl string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			words      []entity.Word
			word       entity.Word
			query      = bson.M{}
			err        error
			returnCode int
		)

		query["word.category"] = "noun"
		query["word.user"] = c.Param("user")

		words, returnCode, err = game.FetchWords(finderUrl, query, 1)
		if err != nil {
			c.JSON(returnCode, fmt.Sprint(err))

			return
		}

		word = words[0]

		game := game.GameOption{}

		game.Question = fmt.Sprintf("What's the article of '%s'?", word.GetGerman())
		game.Option1 = "der"
		game.Option2 = "die"
		game.Option3 = "das"
		game.Id = word.GetId().Hex()

		c.JSON(200, game)
	}
}

func makeCheckAnswerHandle(finderUrl, scorerUrl string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			query      = bson.M{}
			dictionary german.Dictionary
			noun       entity.Noun
			err        error
			returnCode int
		)

		query["__id"] = c.PostForm("id")

		dictionary, returnCode, err = game.FetchDictionary(finderUrl, query, 1)
		if err != nil {
			c.JSON(returnCode, fmt.Sprint(err))

			return
		}

		if len(dictionary.Nouns) == 0 {
			c.JSON(returnCode, "No noun was found.")

			return
		}

		noun = dictionary.Nouns[0]

		answeredRight := checkAnswer(noun, c.PostForm("result"))

		game.ScoreWords(scorerUrl, 10, []string{c.PostForm("id")})

		c.JSON(200, answeredRight)
	}
}

func checkAnswer(word entity.Noun, result string) bool {
	for _, article := range word.Articles {
		if article == entity.Der && result == "1" {
			return true
		}

		if article == entity.Die && result == "2" {
			return true
		}

		if article == entity.Das && result == "3" {
			return true
		}
	}

	return false
}
