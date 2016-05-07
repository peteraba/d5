package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/game/conjugate/lib"
	"github.com/peteraba/d5/lib/game"
	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	game_dbhost_env = "GAME_DBHOST"
	game_dbname_env = "GAME_DBNAME"
)

func parseEnvs() (string, string) {
	// Mongo database host
	hostname := os.Getenv(game_dbhost_env)

	// Mongo database name
	dbName := os.Getenv(game_dbname_env)

	return hostname, dbName
}

func parseFlags() (int, bool, string, string, string) {
	port := flag.Int("port", 17182, "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	finder := flag.String("finder", "http://localhost:17171/", "Finder address")

	scorer := flag.String("scorer", "http://localhost:17172/", "Scorer address")

	collectionName := flag.String("coll", "result", "Collection name for storing results")

	flag.Parse()

	return *port, *debug, *finder, *scorer, *collectionName
}

func main() {
	port, debug, finderUrl, scorerUrl, collectionName := parseFlags()
	hostName, dbName := parseEnvs()

	if hostName == "" || dbName == "" {
		log.Fatalln("Missing environment variables")
	}

	mgoDb, err := mongo.CreateMgoDb(hostName, dbName)
	if err != nil {
		log.Fatalf("MongoDB database could not be created: %s", err)
	}

	mgoCollection := mgoDb.C(collectionName)

	err = mongo.SetResultIndexes(mgoCollection)
	if err != nil {
		log.Println(err)
	}

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/game/:user", makeGameHandle(finderUrl, mgoCollection, debug))
	router.POST("/answer/:user", makeCheckAnswerHandle(finderUrl, scorerUrl, mgoCollection, debug))

	router.Run(fmt.Sprintf(":%d", port))
}

func makeGameHandle(finderUrl string, mgoCollection *mgo.Collection, debug bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var verb *entity.Verb

		query, pp, tense := lib.GetRandomPieces(c.Param("user"))

		dictionary, returnCode, err := game.FetchDictionary(finderUrl, query, 1)
		if err != nil {
			c.JSON(returnCode, fmt.Sprint(err))

			return
		}

		if len(dictionary.Verbs) > 0 {
			verb = &dictionary.Verbs[0]
		}

		gameAnswer, right := getGameAnswer(verb, pp, tense)

		if verb == nil {
			gameAnswer.SetDebugQuery(debug, &query, &dictionary, verb, &bson.M{"pp": pp, "tense": tense})
		} else {
			gameAnswer.SetDebugQuery(debug, &query, nil, verb, &bson.M{"pp": pp, "tense": tense})
			gameAnswer.SetDebugResult(debug, right, verb.GetEnglish(), verb.GetThird())

			err := game.SaveAnswer(gameAnswer.GetId(), []string{verb.GetId().Hex()}, right, mgoCollection)
			if err != nil {
				gameAnswer.Error = fmt.Sprint(err)
			}
		}

		c.JSON(200, gameAnswer)
	}
}

func getGameAnswer(verb *entity.Verb, pp entity.PersonalPronoun, tense entity.Tense) (game.GameAnswer, []string) {
	var right []string

	game := game.GameAnswer{}

	if verb == nil {
		game.Error = "No verbs found"
	} else {
		game.Question = getQuestion(*verb, pp, tense)
		game.Id = util.GenerateUid()

		right = verb.GetVerb(pp, tense)
		if len(right) == 0 {
			game.Error = "No right answer found"
		}
	}

	return game, right
}

func getQuestion(verb entity.Verb, pp entity.PersonalPronoun, tense entity.Tense) string {
	var (
		order      string
		count      string
		tenseLower string
		meanings   []entity.Meaning
		meaning    entity.Meaning
	)

	switch pp {
	case entity.S1:
		order = "1st"
		count = "singular"
		break
	case entity.S2:
		order = "2nd"
		count = "singular"
		break
	case entity.S3:
		order = "3rd"
		count = "singular"
		break
	case entity.P1:
		order = "1st"
		count = "plural"
		break
	case entity.P2:
		order = "2nd"
		count = "plural"
		break
	case entity.P3:
		order = "3rd"
		count = "plural"
		break
	}

	tenseLower = strings.ToLower(fmt.Sprint(tense))

	meanings = verb.GetEnglish()
	if len(meanings) == 0 {
		return fmt.Sprintf("What's the %s person, %s of '%s' in %s tense?", order, count, verb.GetGerman(), tenseLower)
	}

	meaning = meanings[0]

	return fmt.Sprintf("What's the %s person, %s of '%s' in %s tense?", order, count, meaning.Main, tenseLower)
}

func makeCheckAnswerHandle(finderUrl, scorerUrl string, mgoCollection *mgo.Collection, debug bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err        error
			returnCode int
			tracker    game.Tracker
			answerId   string
			answer     string
			score      int
		)

		answerId = c.PostForm("id")
		answer = c.PostForm("answer")

		tracker, err = game.FindAnswer(answerId, mgoCollection)
		if err != nil {
			c.JSON(returnCode, fmt.Sprint(err))

			return
		}

		if util.StringIn(answer, tracker.Right) {
			score = 10
		} else {
			score = 0
		}

		game.ScoreWords(scorerUrl, score, tracker.WordIds)

		c.JSON(200, score)
	}
}
