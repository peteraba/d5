package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	german "github.com/peteraba/d5/lib/german"
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

func parseFlags() (string, string, string, bool) {
	hostName := flag.String("host", "", "Mongo database host")
	dbName := flag.String("db", "", "Mongo database name")
	collectionName := flag.String("coll", "", "Mongo collection for words")
	debug := flag.Bool("debug", false, "Log errors, halt output")

	flag.Parse()

	return *hostName, *dbName, *collectionName, *debug
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

func main() {
	var (
		input        []byte
		err          error
		query        interface{}
		searchResult []german.Superword
		dictionary   german.Dictionary
	)

	hostName, dbName, collectionName, debug := parseFlags()

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
