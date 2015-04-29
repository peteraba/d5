package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/peteraba/d5/shared"
)

func readStdInput(logErrors bool) []byte {
	reader := bufio.NewReader(os.Stdin)

	bytes, err := ioutil.ReadAll(reader)

	if err != nil && logErrors {
		log.Println(err)
	}

	return bytes
}

func parseDictionary(dictionary [][6]string, user string) ([]shared.Word, []string) {
	var (
		words       = []shared.Word{}
		parseErrors = []string{}
	)

	for _, word := range dictionary {
		var (
			w        shared.Word
			german   = word[0]
			english  = word[1]
			third    = word[2]
			category = word[3]
			learned  = word[4]
			score    = word[5]
		)

		if english == "" {
			continue
		}

		switch category {
		case "adj":
			w = shared.NewAdjective(german, english, third, user, learned, score)
			break
		case "noun":
			if shared.NounRegexp.MatchString(german) {
				w = shared.NewNoun(german, english, third, user, learned, score)
			}
			break
		case "verb":
			if shared.VerbRegexp.MatchString(german) {
				w = shared.NewVerb(german, english, third, user, learned, score)
			}
			break
		default:
			w = shared.NewWord(german, english, third, category, user, learned, score, true)
		}

		if w == nil {
			parseErrors = append(parseErrors, german)
			w = shared.NewWord(german, english, third, category, user, learned, score, false)
		}

		words = append(words, w)
	}

	return words, parseErrors
}

func main() {
	var (
		user       = ""
		logErrors  = false
		dictionary = [][6]string{}
	)

	if len(os.Args) > 1 {
		user = os.Args[1]
	}
	if len(os.Args) > 2 {
		logErrors, _ = strconv.ParseBool(os.Args[1])
	}

	json.Unmarshal(readStdInput(logErrors), &dictionary)

	words, parseErrors := parseDictionary(dictionary, user)

	if logErrors && len(parseErrors) > 0 {
		for _, word := range parseErrors {
			log.Printf("Failed: %v\n", word)
		}
	}

	b, err := json.Marshal(words)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
