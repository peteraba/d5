package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/server"
	"github.com/peteraba/d5/lib/util"
)

const name = "persister"
const version = "0.1"
const usage = `
Persister supports CLI and Server mode.

In CLI mode it expects input data on standard input, in server mode as raw POST body

Usage:
  persister [--server] [--port=<n>] [--debug]
  persister -h | --help
  persister -v | --version

Options:
  -s, --server    run in server mode
  -p, --port=<n>  port to open (server mode only) [default: 10020]
  -d, --debug     skip ticks and generate fake data concurrently
  -v, --version   show version information
  -h, --help      show help information

Accepted input data:
  - Raw JSON data to persist

Used environment variables:
  - D5_DBHOST             host or ip of mongodb
  - D5_DBNAME             database name
  - D5_COLLECTION_NAME    collection name
  - D5_COLLECTION_TYPE    collection type
`

/**
 * DOMAIN
 */

func removeUserCollection(collection *mgo.Collection, user string) error {
	if _, err := collection.RemoveAll(bson.M{"word.user": user}); err != nil {
		return err
	}

	return nil
}

func insertWords(collection *mgo.Collection, words []entity.Word) error {
	var (
		err error
	)

	for _, word := range words {
		if word.GetUser() == "" {
			continue
		}

		if err = collection.Insert(word); err != nil {
			return err
		}
	}

	return nil
}

func getPersistResponse(db *mgo.Database, collectionName string, words []entity.Word) error {
	var (
		collection *mgo.Collection
	)

	if len(words) == 0 {
		return errors.New("Words list is empty")
	}

	collection = db.C(collectionName)

	err := removeUserCollection(collection, words[0].GetUser())
	if err != nil {
		return err
	}

	return insertWords(collection, words)
}

/**
 * CLI
 */

func serveCli(mgoDb *mgo.Database, isDebug bool) {
	err := cliHandler(mgoDb, isDebug)

	util.LogFatalErr(err, isDebug)
}

func cliHandler(mgoDb *mgo.Database, isDebug bool) error {
	rawInput, err := util.ReadStdInput()
	if err != nil {
		return err
	}

	words, collectionName, err := getPersistData(rawInput)
	if err != nil {
		return err
	}

	return getPersistResponse(mgoDb, collectionName, words)
}

/**
 * SERVER
 */

func startServer(port int, mgoDb *mgo.Database, isDebug bool) {
	s := server.MakeServer(port, mgoDb, isDebug)

	s.AddHandler("/", persistHandle, server.PostOnly)

	s.Start()
}

func persistHandle(w http.ResponseWriter, r *http.Request, mgoDb *mgo.Database, isDebug bool) error {
	words, collectionName, err := getServerPersistData(r)
	if err != nil {
		return err
	}

	err = getPersistResponse(mgoDb, collectionName, words)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(true)

	return nil
}

func getServerPersistData(r *http.Request) ([]entity.Word, string, error) {
	rawBody, _ := ioutil.ReadAll(r.Body)

	return getPersistData(rawBody)
}

/**
 * INPUT PARSING
 */

func getPersistData(rawInput []byte) ([]entity.Word, string, error) {
	collectionName, _ := mongo.ParseCollectionEnvs()

	words, err := german.ParseWords(rawInput)

	return words, collectionName, err
}

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
