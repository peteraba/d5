package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/game"
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
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
			query      = map[string]string{}
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

		game := game.Game{
			fmt.Sprintf("What's the article of '%s'?", word.GetGerman()),
			"der",
			"die",
			"das",
			"",
			word.GetId().Hex(),
		}

		c.JSON(200, game)
	}
}

func makeCheckAnswerHandle(finderUrl, scorerUrl string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			id   string
			word german.Superword
			err  error
		)

		id = c.PostForm("id")

		data := getCheckAnswerData(id)

		if word, err = findSuperword(finderUrl, data); err != nil {
			c.JSON(500, fmt.Sprintf("%v", err))

			return
		}

		answeredRight := checkAnswer(word, c.PostForm("result"))

		scoreAnswer(scorerUrl, id, answeredRight)

		c.JSON(200, answeredRight)
	}
}

func getCheckAnswerData(id string) url.Values {
	var (
		data  = url.Values{}
		query = map[string]string{}
	)

	query["__id"] = id
	bytes, err := json.Marshal(query)
	if err != nil {
		return data
	}

	data.Set("limit", "1")
	data.Set("query", string(bytes))

	return data
}

func findSuperword(finderUrl string, data url.Values) (german.Superword, error) {
	var (
		err  error
		word = german.Superword{}
	)

	resp, err := http.PostForm(finderUrl, data)
	if err != nil {
		return word, errors.New("Finder call failed.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return word, errors.New("Fetching word by id failed.")
	}

	if err = json.Unmarshal([]byte(body), &word); err != nil {
		return word, errors.New("Parsing word failed.")
	}

	return word, nil
}

func checkAnswer(word german.Superword, result string) bool {
	if word.Articles[0] == entity.Der && result == "1" {
		return true
	}

	if word.Articles[0] == entity.Die && result == "2" {
		return true
	}

	if word.Articles[0] == entity.Das && result == "3" {
		return true
	}

	return false
}

func scoreAnswer(scorerUrl, id string, answeredRight bool) {
	var (
		data  = url.Values{}
		score = "1"
	)

	if answeredRight {
		score = "10"
	}

	data.Set("wordId", id)
	data.Set("score", score)

	http.PostForm(scorerUrl, data)
}
