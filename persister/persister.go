package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/peteraba/d5/shared"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func parseWords(unparsedWords []interface{}) ([]shared.Word, error) {
	var parsedWords = []shared.Word{}

	return parsedWords, nil
}

func removeUserCollection(collection *mgo.Collection, user string) error {
	_, err := collection.RemoveAll(bson.M{"user": user})

	return err
}

func saveCollection(collection *mgo.Collection, words []shared.Word) error {
	var err error

	for _, word := range words {
		if err = collection.Insert(word); err != nil {
			return err
		}
	}

	return nil
}

// Args
//  - program name
//  - mgo url
//  - mgo database
//  - mgo collection
func main() {
	var (
		unparsedWords []interface{}
		parsedWords   []shared.Word
		input         []byte
		err           error
		collection    *mgo.Collection
	)

	if len(os.Args) < 4 {
		log.Fatalln("Mandatory arguments: mgo url, mgo database, mgo collection")
	}

	session, err := mgo.Dial(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	if input, err = readStdInput(); err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(input, &unparsedWords); err != nil {
		log.Fatalln(err)
	}

	parsedWords, err = parseWords(unparsedWords)

	collection = session.DB(os.Args[2]).C(os.Args[3])

	if len(parsedWords) > 0 {
		removeUserCollection(collection, parsedWords[0].GetUser())
	}

	saveCollection(collection, parsedWords)
}
