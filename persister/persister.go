package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	german "github.com/peteraba/d5/lib/german"
	entity "github.com/peteraba/d5/lib/german/entity"
)

const (
	d5_dbhost_env       = "D5_HOSTNAME"
	d5_dbname_env       = "D5_DBNAME"
	d5_coll_words_env   = "D5_COLL_WORDS"
	persister_debug_env = "PERSISTER_DEBUG"
)

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func removeUserCollection(collection *mgo.Collection, user string) error {
	if _, err := collection.RemoveAll(bson.M{"word.user": user}); err != nil {
		return err
	}

	return nil
}

func insertWords(collection *mgo.Collection, words []entity.Word) error {
	var err error

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

func saveCollection(words []entity.Word, url, databaseName, collectionName string) error {
	var (
		collection *mgo.Collection
		err        error
	)

	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection = session.DB(databaseName).C(collectionName)

	if len(words) == 0 {
		return errors.New("Words list is empty")
	}

	removeUserCollection(collection, words[0].GetUser())

	insertWords(collection, words)

	return nil
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
	debugRaw := os.Getenv(persister_debug_env)

	if debugRaw == "1" || strings.ToLower(debugRaw) == "true" {
		debug = true
	}

	return hostname, dbName, collectionName, debug
}

func main() {
	var (
		words []entity.Word
		input []byte
		err   error
	)

	hostName, dbName, collectionName, debug := parseEnvs()

	if input, err = readStdInput(); err != nil {
		log.Fatalln(err)
	}

	words, err = german.ParseWords(input)
	if err != nil {
		log.Fatalln(err)
	}

	err = saveCollection(words, hostName, dbName, collectionName)
	if err != nil {
		log.Fatalln(err)
	}

	if debug {
		log.Printf("Words count: %d\n", len(words))
	}
}
