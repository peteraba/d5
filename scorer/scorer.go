package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/repository"
	"github.com/peteraba/d5/lib/server"
	"github.com/peteraba/d5/lib/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const name = "scorer"
const version = "0.1"
const usage = `
Scorer supports CLI and Server mode.

In CLI mode it expects input data on standard input as JSON, in server mode as a standard form.

Usage:
  scorer [--server] [--port=<n>] [--debug]
  scorer -h | --help
  scorer -v | --version

Options:
  -s, --server    run in server mode
  -p, --port=<n>  port to open (server mode only) [default: 10230]
  -d, --debug     skip ticks and generate fake data concurrently
  -v, --version   show version information
  -h, --help      show help information

Accepted input data:
  - wordId  Word id to find
  - score   Score to associate

Environment variables:
  - D5_DB_HOST                  database host or ip
  - D5_DB_NAME                  database name
  - D5_GAME_TYPE                game type
  - D5_COLLECTION_DATA_GENERAL  name of general collection
  - D5_COLLECTION_DATA_GERMAN   name of german collection
  - D5_COLLECTION_DATA_RESULT   name of result collection
`

/**
 * MAIN
 */

func main() {
	isServer, port, isDebug := util.GetServerOptions(util.GetCliArguments(usage, name, version))

	mgoDb := mongo.CreateMgoDbFromEnvs()

	if isServer {
		startServer(port, mgoDb, isDebug)
		return
	}

	serveCli(mgoDb, isDebug)
}

/**
 * DOMAIN
 */

func getScoreResponse(repo repository.QueryRepo, collectionName string, wordId string, score int) (bool, error) {
	var (
		err      error
		word     entity.Word
		objectId *bson.ObjectId
	)

	objectId = util.HexToObjectId(wordId)
	if objectId == nil {
		return false, errors.New(fmt.Sprintf("Word not found: %s", wordId))
	}

	word, err = repo.FetchWord(collectionName, *objectId)
	if err != nil {
		return false, err
	}

	word.NewScore(score)

	err = repo.UpdateWord(collectionName, *objectId, word)

	return err == nil, err
}

func saveScore(mgoDb *mgo.Database, collectionName string, wordId string, score int) (bool, error) {
	repo := repository.CreateRepo(mgoDb)

	return getScoreResponse(repo, collectionName, wordId, score)
}

/**
 * CLI
 */

func serveCli(mgoDb *mgo.Database, isDebug bool) {
	result, err := cliHandler(mgoDb, isDebug)

	util.LogFatalErr(err, isDebug)

	fmt.Print(result)
}

func cliHandler(mgoDb *mgo.Database, isDebug bool) (interface{}, error) {
	stdInput, err := util.ReadStdInput()
	if err != nil {
		return nil, err
	}

	wordId, score, collectionName, err := getCliScoreData(stdInput)
	if err != nil {
		return nil, err
	}

	data, err := saveScore(mgoDb, collectionName, wordId, score)
	if err != nil {
		return nil, err
	}

	return util.DataToJson(data, isDebug)
}

func getCliScoreData(stdInput []byte) (string, int, string, error) {
	input := strings.Trim(string(stdInput), "\n\t ")
	values, err := url.ParseQuery(input)

	if err != nil {
		return "", 0, "", err
	}

	wordId := values.Get("wordId")
	rawScore := values.Get("score")

	return getScoreData(wordId, rawScore)
}

/**
 * SERVER
 */

func startServer(port int, mgoDb *mgo.Database, isDebug bool) {
	s := server.MakeServer(port, mgoDb, isDebug)

	s.AddHandler("/", scoreHandle, server.PostOnly)

	s.Start()
}

func scoreHandle(w http.ResponseWriter, r *http.Request, mgoDb *mgo.Database, isDebug bool) error {
	wordId, score, collectionName, err := getServerScoreData(r)
	if err != nil {
		return err
	}

	data, err := saveScore(mgoDb, collectionName, wordId, score)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(data)

	return nil
}

func getServerScoreData(r *http.Request) (string, int, string, error) {
	wordId := r.FormValue("wordId")
	score := r.FormValue("score")

	return getScoreData(wordId, score)
}

/**
 * INPUT PARSING
 */

func getScoreData(wordId, rawScore string) (string, int, string, error) {
	collectionName := mongo.ParseDataCollection()

	return filterData(wordId, rawScore, collectionName)
}

func filterData(wordId, rawScore, collectionName string) (string, int, string, error) {
	if wordId == "" {
		return "", 0, "", errors.New("Word id was not posted.")
	}

	if rawScore == "" {
		return "", 0, "", errors.New("Score was not posted.")
	}

	score64, err := strconv.ParseInt(rawScore, 10, 0)
	if err != nil {
		return "", 0, "", errors.New("Score is not a valid integer")
	}

	if score64 < -10 || score64 > 10 {
		return "", 0, "", errors.New("Score is not between -10 and 10")
	}

	return wordId, int(score64), collectionName, nil
}
