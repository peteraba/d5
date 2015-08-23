package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/game"
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	game_dbhost_env = "GAME_HOSTNAME"
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

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/game/:user", makeGameHandle(finderUrl, mgoDb, collectionName, debug))
	router.POST("/answer/:user", makeCheckAnswerHandle(finderUrl, scorerUrl, hostName, dbName, collectionName, debug))

	router.Run(fmt.Sprintf(":%d", port))
}

func makeGameHandle(finderUrl string, mgoDb *mgo.Database, collectionName string, debug bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var verb *entity.Verb

		query, pp, tense := getRandomPieces(c.Param("user"))

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

			err := game.SaveAnswer(&gameAnswer, mgoDb, collectionName)
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

func getRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	switch rand.Int31n(6) {
	case 0:
		return getS2RandomPieces(user)
	case 1:
		return getPastRandomPieces(user)
	case 2:
		return getPastParticleRandomPieces(user)
	}

	return getGeneralRandomPieces(user)
}

func getS2RandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		query = getBaseQuery(user)
	)

	query["s2"] = bson.M{"$regex": bson.RegEx{".*", "s"}}

	switch rand.Int31n(2) {
	case 0:
		pp = entity.S2
		break
	case 1:
		pp = entity.S3
		break
	}

	return query, pp, entity.Present
}

func getPastRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		tense entity.Tense
		query bson.M
	)

	query, pp, tense = getGeneralRandomPieces(user)

	query["preterite"] = bson.M{"$regex": bson.RegEx{".*", "s"}}

	switch rand.Int31n(3) {
	case 0:
	case 1:
		tense = entity.PastParticiple
		break
	case 2:
		tense = entity.Preterite
		break
	}

	return query, pp, tense
}

func getPastParticleRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		query = getBaseQuery(user)
	)

	query, pp, _ = getGeneralRandomPieces(user)

	query["auxiliary"] = "s"

	return query, pp, entity.PastParticiple
}

func getGeneralRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		tense entity.Tense
		query = getBaseQuery(user)
	)

	switch rand.Int31n(3) {
	case 0:
		tense = entity.Present
		break
	case 1:
		tense = entity.Preterite
	case 2:
		tense = entity.PastParticiple
	}

	switch rand.Int31n(6) {
	case 0:
		return query, entity.S1, tense
	case 1:
		return query, entity.S2, tense
	case 2:
		return query, entity.S3, tense
	case 3:
		return query, entity.P1, tense
	case 4:
		return query, entity.P2, tense
	}

	return query, entity.P3, tense
}

func getBaseQuery(user string) bson.M {
	var (
		query = bson.M{}
	)

	query["word.category"] = "verb"
	query["word.user"] = user

	return query
}

func getQuestion(verb entity.Verb, pp entity.PersonalPronoun, tense entity.Tense) string {
	var (
		order      string
		count      string
		tenseLower string
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

	return fmt.Sprintf("What's the %s person, %s of '%s' in %s tense?", order, count, verb.GetGerman(), tenseLower)
}

func makeCheckAnswerHandle(finderUrl, scorerUrl, hostName, dbName, collectionName string, debug bool) func(c *gin.Context) {
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
