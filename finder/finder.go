package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/repository"
	"github.com/peteraba/d5/lib/server"
	"github.com/peteraba/d5/lib/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const name = "finder"
const version = "0.1"
const usage = `
Finder supports CLI and Server mode.

In CLI mode it expects input data on standard input as JSON, in server mode as a standard form.

Usage:
  finder [--server] [--port=<n>] [--debug]
  finder -h | --help
  finder -v | --version

Options:
  -s, --server    run in server mode
  -p, --port=<n>  port to open (server mode only) [default: 10110]
  -d, --debug     skip ticks and generate fake data concurrently
  -v, --version   show version information
  -h, --help      show help information

Accepted input data:
  - query  Search query as JSON string
  - limit  Maximum number of items to be returned [default: 100]

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

	mgoDb, err := mongo.CreateMgoDbFromEnvs()
	util.LogFatalfMsg(err, "MongoDB database could not be created: %v", true)

	if isServer {
		startServer(port, mgoDb, isDebug)
	} else {
		serveCli(mgoDb, isDebug)
	}
}

/**
 * DOMAIN
 */

func getFinderResponse(repo repository.QueryRepo, collectionName string, query bson.M, limit int) (interface{}, error) {
	var (
		objectId *bson.ObjectId
		word     entity.Word
		err      error
	)

	if _, ok := query["__id"]; ok {
		objectId = util.HexToObjectId(query["__id"].(string))
		if objectId == nil {
			return nil, errors.New(fmt.Sprintf("String can not be turned into hex: %v", objectId))
		}

		word, err = repo.FetchWord(collectionName, *objectId)

		return []entity.Word{word}, nil
	}

	if _, ok := query["word.user"]; !ok {
		return nil, errors.New("word.user key must be defined for searches.")
	}

	_, err = repo.FetchDictionary(collectionName, query)
	if err != nil {
		return nil, err
	}

	return repo.FilterDictionary(limit)
}

func findWords(mgoDb *mgo.Database, collectionName string, query bson.M, limit int) (interface{}, error) {
	repo := repository.CreateRepo(mgoDb)

	return getFinderResponse(repo, collectionName, query, limit)
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

	query, limit, collectionName, err := getCliFinderData(stdInput)
	if err != nil {
		return nil, err
	}

	data, err := findWords(mgoDb, collectionName, query, limit)
	if err != nil {
		return nil, err
	}

	return util.DataToJson(data, isDebug)
}

func getCliFinderData(stdInput []byte) (bson.M, int, string, error) {
	values, err := url.ParseQuery(string(stdInput))

	if err != nil {
		return bson.M{}, 0, "", err
	}

	rawQuery := values.Get("query")
	rawLimit := values.Get("limit")

	return getFinderData(rawQuery, rawLimit)
}

/**
 * SERVER
 */

func startServer(port int, mgoDb *mgo.Database, isDebug bool) {
	s := server.MakeServer(port, mgoDb, isDebug)

	s.AddHandler("/", findHandle, server.PostOnly)

	s.Start()
}

func findHandle(w http.ResponseWriter, r *http.Request, mgoDb *mgo.Database, isDebug bool) error {
	query, limit, collectionName, err := getServerFinderData(r)
	if err != nil {
		return err
	}

	data, err := findWords(mgoDb, collectionName, query, limit)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(data)

	return nil
}

func getServerFinderData(r *http.Request) (bson.M, int, string, error) {
	rawQuery := r.FormValue("query")
	rawLimit := r.FormValue("limit")

	return getFinderData(rawQuery, rawLimit)
}

/**
 * INPUT PARSING
 */

func getFinderData(rawQuery, rawLimit string) (bson.M, int, string, error) {
	collectionName := mongo.ParseDataCollection()

	return filterData(rawQuery, rawLimit, collectionName)
}

func filterData(rawQuery, rawLimit, collectionName string) (bson.M, int, string, error) {
	search := bson.M{}

	if rawQuery == "" {
		return search, 0, "", errors.New("Query was not posted.")
	}

	if rawLimit == "" {
		return search, 0, "", errors.New("Limit was not posted.")
	}

	limit64, err := strconv.ParseInt(rawLimit, 10, 0)
	if err != nil {
		return search, 0, "", errors.New("Limit is not a valid integer")
	}

	if limit64 < 0 {
		return search, 0, "", errors.New("Limit is not smaller than 0")
	}

	err = json.Unmarshal([]byte(rawQuery), &search)

	return search, int(limit64), collectionName, err
}
