package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/game/conjugate/lib"
	game "github.com/peteraba/d5/game/lib"
	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const name = "Conjugate"
const version = "0.1"
const defaultPort = "10420"

/**
 * MAIN
 */

func main() {
	gameServer := Conjugate{}
	game.Main(name, version, defaultPort, gameServer)
}

type Conjugate struct{}

func (conjugate Conjugate) MakeGameHandle(finderUrl string, mgoCollection *mgo.Collection, debug bool) func(context *gin.Context) {
	return func(context *gin.Context) {
		var verb *entity.Verb

		query, pp, tense := lib.GetRandomPieces(context.Param("user"))

		dictionary, returnCode, err := game.FetchDictionary(finderUrl, query, 1)
		if err != nil {
			context.JSON(returnCode, fmt.Sprint(err))

			return
		}

		if len(dictionary.Verbs) > 0 {
			verb = &dictionary.Verbs[0]
		}

		gameAnswer, isRight := conjugate.getGameAnswer(verb, pp, tense)

		if verb == nil {
			gameAnswer.SetDebugQuery(debug, &query, &dictionary, verb, &bson.M{"pp": pp, "tense": tense})
		} else {
			gameAnswer.SetDebugQuery(debug, &query, nil, verb, &bson.M{"pp": pp, "tense": tense})
			gameAnswer.SetDebugResult(debug, isRight, verb.GetEnglish(), verb.GetThird())

			err := game.SaveAnswer(gameAnswer.GetId(), []string{verb.GetId().Hex()}, isRight, mgoCollection)
			if err != nil {
				gameAnswer.Error = fmt.Sprint(err)
			}
		}

		context.JSON(200, gameAnswer)
	}
}

func (conjugate Conjugate) getGameAnswer(verb *entity.Verb, pp entity.PersonalPronoun, tense entity.Tense) (game.GameAnswer, []string) {
	var right []string

	game := game.GameAnswer{}

	if verb == nil {
		game.Error = "No verbs found"
	} else {
		game.Question = conjugate.getQuestion(*verb, pp, tense)
		game.Id = util.GenerateUid()

		right = verb.GetVerb(pp, tense)
		if len(right) == 0 {
			game.Error = "No right answer found"
		}
	}

	return game, right
}

func (conjugate Conjugate) getQuestion(verb entity.Verb, pp entity.PersonalPronoun, tense entity.Tense) string {
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

func (conjugate Conjugate) MakeCheckAnswerHandle(finderUrl, scorerUrl string, mgoCollection *mgo.Collection, debug bool) func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			err        error
			returnCode int
			tracker    game.Tracker
			answerId   string
			answer     string
			score      int
		)

		answerId = context.PostForm("id")
		answer = context.PostForm("answer")

		tracker, err = game.FindAnswer(answerId, mgoCollection)
		if err != nil {
			context.JSON(returnCode, fmt.Sprint(err))

			return
		}

		if util.StringIn(answer, tracker.Right) {
			score = 10
		} else {
			score = 0
		}

		game.ScoreWords(scorerUrl, score, tracker.WordIds)

		context.JSON(200, score)
	}
}
