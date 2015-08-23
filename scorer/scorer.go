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

func getSearchQuery(bytes []byte) (map[string]string, error) {
	var search = make(map[string]string)

	err := json.Unmarshal(bytes, &search)

	return search, err
}

/**
 * DOMAIN
 */

func getResponseData(repo repository.QueryRepo, collectionName string, wordId string, score int) (bool, error) {
	var (
		err      error
		word     entity.Word
		objectId *bson.ObjectId
	)

	objectId = util.HexToObjectId(wordId)

	word, err = repo.FetchWord(collectionName, *objectId)
	if err != nil {
		return false, err
	}

	word.NewScore(score)

	err = repo.UpdateWord(collectionName, *objectId, word)
	if err != nil {
		return false, err
	}

	return true, nil
}

/**
 * CLI
 */

func cli(
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
	wordId string,
	score int,
) {
	result, err := cliWrapped(mgoDb, collectionName, isGerman, debug, wordId, score)
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
	wordId string,
	score int,
) (interface{}, error) {
	var (
		err  error
		data interface{}
	)

	repo := repository.CreateRepo(mgoDb, isGerman)

	data, err = getResponseData(repo, collectionName, wordId, score)
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

func server(port int, mgoDb *mgo.Database, collectionName string, isGerman bool, debug bool) {
	http.HandleFunc("/", makeHandler(scoreHandle, mgoDb, collectionName, isGerman, debug))

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

func scoreHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) error {
	wordId, score, err := getUpdateData(r)
	if err != nil {
		return err
	}

	repo := repository.CreateRepo(mgoDb, isGerman)

	data, err := getResponseData(repo, collectionName, wordId, score)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(data)

	return nil
}

func getUpdateData(r *http.Request) (string, int, error) {
	var (
		rawId    string
		rawScore string
		score    int64
		err      error
	)

	r.ParseMultipartForm(1024 * 1024 * 10)

	rawId = r.Form.Get("wordId")
	if rawId == "" {
		return "", 0, errors.New("Word id was not posted.")
	}

	rawScore = r.Form.Get("score")
	if rawScore == "" {
		return "", 0, errors.New("Score was not posted.")
	}

	score, err = strconv.ParseInt(rawScore, 10, 0)
	if err != nil {
		return "", 0, errors.New("Score is not valid integer")
	}

	if score < -10 || score > 10 {
		return "", 0, errors.New("Score is not between -10 and 10")
	}

	return rawId, int(score), nil
}

/**
 * INPUT PARSING
 */

func parseFlags() (bool, int, string, string, bool, string, int) {
	isServer := flag.Bool("server", false, "Starts a server")
	port := flag.Int("port", 17172, "Port for server")

	collectionName := flag.String("coll", "german", "Port for server")
	collectionType := flag.String("type", COLL_TYPE_GERMAN, "Type of collection (german, anything else)")

	debug := flag.Bool("debug", false, "Enables debug logs")

	wordId := flag.String("wordId", "", "Id of word to update")
	score := flag.Int("score", 0, "Score")

	flag.Parse()

	return *isServer, *port, *collectionName, *collectionType, *debug, *wordId, *score
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
		log.Fatalf("MongoDB database could not be created: %v", err)
	}

	isServer, port, collectionName, collectionType, debug, wordId, score := parseFlags()

	isGerman := !(collectionType == "" || collectionType == COLL_TYPE_DEFAULT)

	if isServer {
		server(port, mgoDb, collectionName, isGerman, debug)
	} else {
		cli(mgoDb, collectionName, isGerman, debug, wordId, score)
	}
}
