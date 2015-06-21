package main

import (
	"bufio"
	"encoding/json"
	"flag"
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

func runQuery(query interface{}, hostName, dbName, collectionName string) german.Dictionary {
	var (
		err          error
		searchResult []german.Superword
	)

	searchResult, err = getCollection(query, hostName, dbName, collectionName)
	if err != nil {
		log.Fatalln(err)
	}

	dictionary := german.SuperwordsToDictionary(searchResult)
	if err != nil {
		log.Fatalln(err)
	}

	return dictionary
}

func cli(hostName, dbName, collectionName string, debug bool) {
	var (
		input      []byte
		query      interface{}
		err        error
		dictionary german.Dictionary
	)

	if input, err = readStdInput(); err != nil {
		log.Fatalln(err)
	}

	query, err = getSearchQuery(input)
	if err != nil {
		log.Fatalln(err)
	}

	dictionary = runQuery(query, hostName, dbName, collectionName)

	outputJson(dictionary, debug)
}

func server(port int, hostName, dbName, collectionName string, debug bool) {
	log.Print("Server is not yet implemented")
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

func parseFlags() (bool, int) {
	isServer := flag.Bool("server", false, "Starts a server")
	port := flag.Int("port", 17171, "Port for server")

	flag.Parse()

	return *isServer, *port
}

func main() {
	hostName, dbName, collectionName, debug := parseEnvs()

	isServer, port := parseFlags()

	if isServer {
		server(port, hostName, dbName, collectionName, debug)
	} else {
		cli(hostName, dbName, collectionName, debug)
	}
}
