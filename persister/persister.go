package main

import (
	"bufio"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	german "github.com/peteraba/d5/lib/german"
	entity "github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/mongo"
)

const (
	dbhost_env = "D5_DBHOST"
	dbname_env = "D5_DBNAME"
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

func saveCollection(words []entity.Word, db *mgo.Database, collectionName string) error {
	var (
		collection *mgo.Collection
	)

	if len(words) == 0 {
		return errors.New("Words list is empty")
	}

	collection = db.C(collectionName)

	removeUserCollection(collection, words[0].GetUser())

	insertWords(collection, words)

	return nil
}

func parseEnvs() (string, string) {
	// Mongo database host
	hostname := os.Getenv(dbhost_env)

	// Mongo database name
	dbName := os.Getenv(dbname_env)

	return hostname, dbName
}

func parseFlags() (string, bool) {
	collectionName := flag.String("coll", "german", "Port for server")

	debug := flag.Bool("debug", false, "Enables debug logs")

	flag.Parse()

	return *collectionName, *debug
}

func main() {
	var (
		words []entity.Word
		input []byte
		err   error
	)

	hostName, dbName := parseEnvs()
	if hostName == "" || dbName == "" {
		log.Fatalln("Missing environment variables")
	}

	mgoDb, err := mongo.CreateMgoDb(hostName, dbName)
	if err != nil {
		log.Fatalf("MongoDB database could not be created: %s", err)
	}

	collectionName, debug := parseFlags()

	if input, err = readStdInput(); err != nil {
		log.Fatalln(err)
	}

	words, err = german.ParseWords(input)
	if err != nil {
		log.Fatalln(err)
	}

	err = saveCollection(words, mgoDb, collectionName)
	if err != nil {
		log.Fatalln(err)
	}

	if debug {
		log.Printf("Words count: %d\n", len(words))
	}
}
