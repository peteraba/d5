package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
)

func FetchDictionary(finderUrl string, query bson.M, limit int) (german.Dictionary, int, error) {
	var (
		dictionary german.Dictionary
		err        error
		returnCode int
		body       []byte
	)

	body, returnCode, err = retrieveWords(finderUrl, query, limit)
	if err != nil {
		return dictionary, returnCode, err
	}

	dictionary, err = german.ParseDictionary(body)
	if err != nil {
		log.Printf("Parsing finder response error: %v\n%v\n", err, string(body))

		return dictionary, 500, errors.New("Parsing finder response failed.")
	}

	return dictionary, 200, nil

}

func FetchWords(finderUrl string, query bson.M, limit int) ([]entity.Word, int, error) {
	var (
		words      = []entity.Word{}
		err        error
		returnCode int
		body       []byte
	)

	body, returnCode, err = retrieveWords(finderUrl, query, limit)
	if err != nil {
		return words, returnCode, err
	}

	words, err = german.ParseWords(body)
	if err != nil {
		log.Printf("Parsing finder response error: %v\n", err)

		return words, 500, errors.New("Parsing finder response failed.")
	}

	if len(words) < 1 {
		return words, 204, errors.New("No words returned.")
	}

	return words, 200, nil

}

func retrieveWords(finderUrl string, query bson.M, limit int) ([]byte, int, error) {
	var (
		data  = url.Values{}
		bytes []byte
		err   error
	)

	bytes, err = json.Marshal(query)

	data.Set("limit", fmt.Sprintf("%d", limit))
	data.Set("query", string(bytes))

	resp, err := http.PostForm(finderUrl, data)
	if err != nil {
		log.Printf("Finder call error: %v\n", err)

		return bytes, 503, errors.New("Finder call failed.")
	}

	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading finder response error: %v\n", err)

		return bytes, 500, errors.New("Reading finder response failed.")
	}

	return bytes, 200, nil
}

func ScoreWords(scorerUrl string, score int, ids []string) {
	var (
		data = url.Values{}
	)

	for i := 0; i < len(ids); i++ {
		data.Set("wordId", ids[i])
		data.Set("score", fmt.Sprintf("%d", score))

		http.PostForm(scorerUrl, data)
	}
}
