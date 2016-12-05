package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	game "github.com/peteraba/d5/game/lib"
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const name = "DerDieDas"
const version = "0.1"
const defaultPort = "10410"

type DerDieDas struct{}

/**
 * MAIN
 */

func main() {
	gameServer := DerDieDas{}
	game.Main(name, version, defaultPort, gameServer)
}

func (d DerDieDas) MakeGameHandle(finderUrl string, mgoCollection *mgo.Collection, debug bool) func(c *gin.Context) {
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

func (d DerDieDas) MakeCheckAnswerHandle(finderUrl, scorerUrl string, mgoCollection *mgo.Collection, isDebug bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			query       = bson.M{}
			dictionary  german.Dictionary
			noun        entity.Noun
			err         error
			returnCode  int
			answerScore int
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

		answerScore = d.CheckAnswer(noun, c.PostForm("answer"))

		game.ScoreWords(scorerUrl, answerScore, []string{c.PostForm("id")})

		c.JSON(200, answerScore)
	}
}

func (d DerDieDas) CheckAnswer(word entity.Noun, result string) int {
	for _, article := range word.Articles {
		if article == entity.Der && result == "1" {
			return 10
		}

		if article == entity.Die && result == "2" {
			return 10
		}

		if article == entity.Das && result == "3" {
			return 10
		}
	}

	return 0
}
