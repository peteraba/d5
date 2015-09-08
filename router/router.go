package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	if objectId == nil {
		return false, errors.New(fmt.Sprintf("WordId could not be converted: %s", wordId))
	}

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
 * SERVER
 */
func server(port int, mgoDb *mgo.Database, collectionName string, debug bool) {
	http.HandleFunc("/game/", makeHandler(gameHandle, mgoDb, collectionName, debug))
	http.HandleFunc("/answer/", makeHandler(answerHandle, mgoDb, collectionName, debug))

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func makeHandler(
	fn func(http.ResponseWriter, *http.Request, *mgo.Database, string, bool) error,
	mgoDb *mgo.Database,
	collectionName string,
	debug bool,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r, mgoDb, collectionName, debug)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprint(err))
			log.Println(err)
		}
	}
}

func gameHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoDb *mgo.Database,
	collectionName string,
	debug bool,
) error {
	user := r.URL.Path[6:]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.Get(findUrl("game", user))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(body))

	return nil
}

func answerHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoDb *mgo.Database,
	collectionName string,
	debug bool,
) error {
	user := r.URL.Path[6:]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.PostForm(findUrl("answer", user), r.Form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(body))

	return nil
}

func findUrl(path, user string) string {

	return "http://localhost:1782/" + path + "/" + user
}

func getAnswerData(r *http.Request) (string, string, string) {
	var (
		user   string
		id     string
		answer string
	)

	user = r.Form.Get("user")

	id = r.PostForm.Get("id")

	answer = r.PostForm.Get("answer")

	return user, id, answer
}

/**
 * INPUT PARSING
 */
func parseFlags() (int, string, bool) {
	port := flag.Int("port", 17173, "Port for server")

	collectionName := flag.String("coll", "german", "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *port, *collectionName, *debug
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

	port, collectionName, debug := parseFlags()

	server(port, mgoDb, collectionName, debug)
}
