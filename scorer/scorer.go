package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/repository"
	"github.com/peteraba/d5/lib/util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

	util.LogFatalErr(err, debug)

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

	return util.DataToJson(data, debug)
}

/**
 * SERVER
 */

func server(port int, mgoDb *mgo.Database, collectionName string, isGerman bool, debug bool) {
	s := util.MakeServer(port, mgoDb, collectionName, isGerman, debug)

	s.AddHandler("/", scoreHandle, util.PostOnly)

	s.Start()
}

func scoreHandle(
	w http.ResponseWriter,
	r *http.Request,
	mgoDb *mgo.Database,
	collectionName string,
	isGerman bool,
	debug bool,
) error {
	wordId, score, err := getFormData(r)
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

func getFormData(r *http.Request) (string, int, error) {
	var (
		rawId    string
		rawScore string
		score    int64
		err      error
	)

	r.ParseForm()

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
	isServer, port, collectionName, collectionType, debug, data := util.ParseFlags()

	wordId, _ := data["wordId"]

	tmp, _ := data["score"]
	score, _ := strconv.ParseInt(tmp, 10, 64)

	return isServer, port, collectionName, collectionType, debug, wordId, int(score)
}

/**
 * MAIN
 */

func main() {
	hostName, dbName := util.ParseEnvs()
	if hostName == "" || dbName == "" {
		util.LogMsg("Missing environment variables", true, true)
	}

	mgoDb, err := mongo.CreateMgoDb(hostName, dbName)
	util.LogFatalfMsg(err, "MongoDB database could not be created: %v", true)

	isServer, port, collectionName, collectionType, debug, wordId, score := parseFlags()

	isGerman := util.IsGerman(collectionType)

	if isServer {
		server(port, mgoDb, collectionName, isGerman, debug)
	} else {
		cli(mgoDb, collectionName, isGerman, debug, wordId, score)
	}
}
