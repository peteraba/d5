package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/peteraba/d5/finder/repository"
	"github.com/peteraba/d5/lib/mongo"
	"gopkg.in/mgo.v2"
)

const (
	COLL_TYPE_DEFAULT = "default"
	COLL_TYPE_GERMAN  = "german"
)

/**
 * MGO
 */

func getSearchQuery(bytes []byte) (map[string]string, error) {
	var search = make(map[string]string)

	err := json.Unmarshal(bytes, &search)

	return search, err
}

/**
 * DOMAIN
 */

func getResponseData(repo repository.QueryRepo, collectionName string, query map[string]string) (interface{}, error) {
	if _, ok := query["word.user"]; !ok {
		return nil, errors.New("word.user key must be defined for searches.")
	}

	return repo.CreateDictionary(collectionName, query)
}

/**
 * CLI
 */

func cli(
	mgoSession *mgo.Session,
	dbName,
	collectionName string,
	isGerman bool,
	debug bool,
) {
	result, err := cliWrapped(mgoSession, dbName, collectionName, isGerman, debug)
	if err != nil {
		if debug {
			log.Println(err)
		}

		return
	}

	fmt.Print(result)
}

func cliWrapped(
	mgoSession *mgo.Session,
	dbName,
	collectionName string,
	isGerman bool,
	debug bool,
) (interface{}, error) {
	var (
		input []byte
		query map[string]string
		err   error
		data  interface{}
	)

	if input, err = readStdInput(); err != nil {
		return nil, err
	}

	query, err = getSearchQuery(input)
	if err != nil {
		return nil, err
	}

	repo := repository.CreateRepo(mgoSession, dbName, isGerman)

	data, err = getResponseData(repo, collectionName, query)
	if err != nil {
		return nil, err
	}

	return dataToJson(data, debug)
}

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func dataToJson(rawData interface{}, debug bool) (string, error) {
	var (
		bytes []byte
		err   error
	)

	if debug {
		bytes, err = json.MarshalIndent(rawData, "", "  ")
	} else {
		bytes, err = json.Marshal(rawData)
	}

	return fmt.Sprintf("%s\n", string(bytes)), err
}

/**
 * SERVER
 */

func server(port int, mgoSession *mgo.Session, dbName, collectionName string, isGerman bool, debug bool) {
	http.HandleFunc("/", makeHandler(findHandle, mgoSession, dbName, collectionName, isGerman, debug))

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func makeHandler(
	fn func(http.ResponseWriter, *http.Request, *mgo.Session, string, string, bool, bool) error,
	mgoSession *mgo.Session,
	dbName string,
	collectionName string,
	isGerman bool,
	debug bool,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)

			return
		}

		err := fn(w, r, mgoSession, dbName, collectionName, isGerman, debug)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprint(err))
			log.Println(err)
		}
	}
}

func findHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoSession *mgo.Session,
	dbName,
	collectionName string,
	isGerman bool,
	debug bool,
) error {
	rawQuery, err := getQueryValue(r)
	if err != nil {
		return err
	}

	query, err := getSearchQuery([]byte(rawQuery))
	if err != nil {
		return err
	}

	repo := repository.CreateRepo(mgoSession.Clone(), dbName, isGerman)

	data, err := getResponseData(repo, collectionName, query)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(data)

	return nil
}

func getQueryValue(r *http.Request) (string, error) {
	var (
		rawQuery string
	)

	r.ParseMultipartForm(1024 * 1024 * 10)

	rawQuery = r.Form.Get("query")
	if rawQuery == "" {
		return "", errors.New("Query was not posted.")
	}

	return rawQuery, nil
}

/**
 * INPUT PARSING
 */

func parseFlags() (bool, int, string, string, bool) {
	isServer := flag.Bool("server", false, "Starts a server")
	port := flag.Int("port", 17171, "Port for server")

	collectionName := flag.String("coll", "german", "Port for server")
	collectionType := flag.String("type", COLL_TYPE_GERMAN, "Type of collection (german, anything else)")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *isServer, *port, *collectionName, *collectionType, *debug
}

/**
 * MAIN
 */

func main() {
	hostName, dbName := mongo.ParseEnvs()

	isServer, port, collectionName, collectionType, debug := parseFlags()

	isGerman := !(collectionType == "" || collectionType == COLL_TYPE_DEFAULT)

	mgoSession, err := mongo.GetMgoSession(hostName)
	if err != nil {
		log.Fatalf("MongoDB session could not be built")
	}

	if isServer {
		server(port, mgoSession, dbName, collectionName, isGerman, debug)
	} else {
		cli(mgoSession, dbName, collectionName, isGerman, debug)
	}
}
