package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/mgo.v2"

	german "github.com/peteraba/d5/lib/german"
)

const (
	d5_dbhost_env     = "D5_HOSTNAME"
	d5_dbname_env     = "D5_DBNAME"
	d5_coll_words_env = "D5_COLL_WORDS"
	finder_debug_env  = "FINDER_DEBUG"
)

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func getCollection(query interface{}, url, databaseName, collectionName string) ([]german.Superword, error) {
	var (
		collection *mgo.Collection
		err        error
		result     = []german.Superword{}
	)

	session, err := mgo.Dial(url)
	if err != nil {
		return result, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{})

	collection = session.DB(databaseName).C(collectionName)

	err = collection.Find(query).All(&result)

	return result, err
}

func getSearchQuery(bytes []byte) (interface{}, error) {
	var search = make(map[string]string)

	err := json.Unmarshal(bytes, &search)

	return search, err
}

func outputJson(rawData interface{}, debug bool) {
	var (
		bytes []byte
		err   error
	)

	if debug {
		bytes, err = json.MarshalIndent(rawData, "", "  ")
	} else {
		bytes, err = json.Marshal(rawData)
	}

	if err == nil {
		fmt.Printf("%s\n", string(bytes))
	} else if debug {
		log.Println("Data could not be parsed")
	}
}

func parseEnvs() (string, string, string, bool) {
	var debug = false

	// Mongo database host
	hostname := os.Getenv(d5_dbhost_env)

	// Mongo database name
	dbName := os.Getenv(d5_dbname_env)

	// Mongo collection name
	collectionName := os.Getenv(d5_coll_words_env)

	// Is debugging enabled
	debugRaw := os.Getenv(finder_debug_env)

	if debugRaw == "1" || strings.ToLower(debugRaw) == "true" {
		debug = true
	}

	return hostname, dbName, collectionName, debug
}

func main() {
	var (
		input        []byte
		err          error
		query        interface{}
		searchResult []german.Superword
		dictionary   german.Dictionary
	)

	hostName, dbName, collectionName, debug := parseEnvs()

	if input, err = readStdInput(); err != nil {
		log.Fatalln(err)
	}

	query, err = getSearchQuery(input)
	if err != nil {
		log.Fatalln(err)
	}

	searchResult, err = getCollection(query, hostName, dbName, collectionName)
	if err != nil {
		log.Fatalln(err)
	}

	dictionary = german.SuperwordsToDictionary(searchResult)
	if err != nil {
		log.Fatalln(err)
	}

	outputJson(dictionary, debug)
}
