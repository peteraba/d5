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
)

func FetchWords(finderUrl string, query map[string]string, limit int) ([]entity.Word, int, error) {
	var (
		data  = url.Values{}
		bytes []byte
		err   error
		words = []entity.Word{}
	)

	bytes, err = json.Marshal(query)

	data.Set("limit", fmt.Sprintf("%d", limit))
	data.Set("query", string(bytes))

	resp, err := http.PostForm(finderUrl, data)
	if err != nil {
		log.Printf("Finder call error: %v\n", err)

		return words, 503, errors.New("Finder call failed.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading finder response error: %v\n", err)

		return words, 500, errors.New("Reading finder response failed.")
	}

	words, err = german.ParseWords(body)
	if err != nil {
		log.Printf("Parsing finder response error: %v\n", err)

		return words, 500, errors.New("Parsing finder response failed.")
	}

	if len(words) < 0 {
		return words, 204, errors.New("No words returned.")
	}

	return words, 200, nil
}
