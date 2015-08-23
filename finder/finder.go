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
	"net/url"
	"os"
	"strconv"

	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/repository"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	d5_dbhost_env = "D5_DBHOST"
	d5_dbname_env = "D5_DBNAME"
)

const (
	COLL_TYPE_DEFAULT = "default"
	COLL_TYPE_GERMAN  = "german"
)

/**
 * MGO
 */

func getSearchQuery(rawQuery string) (bson.M, error) {
	var search = bson.M{}

	err := json.Unmarshal([]byte(rawQuery), &search)

	return search, err
}

/**
 * DOMAIN
 */

func getResponseData(repo repository.QueryRepo, collectionName string, query bson.M, limit int) (interface{}, error) {
	var (
		objectId *bson.ObjectId
		word     entity.Word
		err      error
	)

	if _, ok := query["__id"]; ok {
		objectId = util.HexToObjectId(query["__id"].(string))

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

/**
 * CLI
 */

func cli(
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) {
	result, err := cliWrapped(mgoDb, collectionName, isGerman, debug)
	if err != nil {
		if debug {
			log.Println(err)
		}

		return
	}

	fmt.Print(result)
}

func cliWrapped(
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) (interface{}, error) {
	var (
		query    bson.M
		err      error
		data     interface{}
		rawQuery string
		limit    int
	)

	if rawQuery, limit, err = readStdInput(); err != nil {
		return nil, err
	}

	query, err = getSearchQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	repo := repository.CreateRepo(mgoDb, isGerman)

	data, err = getResponseData(repo, collectionName, query, limit)
	if err != nil {
		return nil, err
	}

	return dataToJson(data, debug)
}

func readStdInput() (string, int, error) {
	var (
		bytes    []byte
		rawQuery string
		limit    int64
		values   url.Values
		err      error
	)

	reader := bufio.NewReader(os.Stdin)

	bytes, err = ioutil.ReadAll(reader)
	if err != nil {
		return "", 0, err
	}

	values, err = url.ParseQuery(string(bytes))

	rawQuery = values.Get("query")

	limit, err = strconv.ParseInt(values.Get("limit"), 10, 0)
	if err != nil {
		return rawQuery, 0, err
	}

	return rawQuery, int(limit), nil
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

func server(port int, mgoDb *mgo.Database, collectionName string, isGerman bool, debug bool) {
	http.HandleFunc("/", makeHandler(findHandle, mgoDb, collectionName, isGerman, debug))

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func makeHandler(
	fn func(http.ResponseWriter, *http.Request, *mgo.Database, string, bool, bool) error,
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)

			return
		}

		err := fn(w, r, mgoDb, collectionName, isGerman, debug)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprint(err))
			log.Println(err)
		}
	}
}

func findHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) error {
	rawQuery, limit, err := getRequestData(r)
	if err != nil {
		return err
	}

	query, err := getSearchQuery(rawQuery)
	if err != nil {
		return err
	}

	repo := repository.CreateRepo(mgoDb, isGerman)

	data, err := getResponseData(repo, collectionName, query, limit)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(data)

	return nil
}

func getRequestData(r *http.Request) (string, int, error) {
	var (
		rawQuery string
		rawLimit string
		limit    int64
		err      error
	)

	r.ParseMultipartForm(1024 * 1024 * 10)

	rawQuery = r.Form.Get("query")
	if rawQuery == "" {
		return "", 0, errors.New("Query was not posted.")
	}

	rawLimit = r.Form.Get("limit")
	limit, err = strconv.ParseInt(rawLimit, 10, 0)
	if err != nil {
		return rawQuery, 0, errors.New("Limit is not valid integer")
	}

	if limit < 0 {
		return rawQuery, 0, errors.New("Limit is less than 0")
	}

	return rawQuery, int(limit), nil
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

func parseEnvs() (string, string) {
	// Mongo database host
	hostname := os.Getenv(d5_dbhost_env)

	// Mongo database name
	dbName := os.Getenv(d5_dbname_env)

	return hostname, dbName
}

/**
 * MAIN
 */

func main() {
	hostName, dbName := parseEnvs()
	if hostName == "" || dbName == "" {
		log.Fatalln("Missing environment variables")
	}

	mgoDb, err := mongo.CreateMgoDb(hostName, dbName)
	if err != nil {
		log.Fatalf("MongoDB database could not be created. %v", err)
	}

	isServer, port, collectionName, collectionType, debug := parseFlags()

	isGerman := !(collectionType == "" || collectionType == COLL_TYPE_DEFAULT)

	if isServer {
		server(port, mgoDb, collectionName, isGerman, debug)
	} else {
		cli(mgoDb, collectionName, isGerman, debug)
	}
}
