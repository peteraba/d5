package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	germanLib "github.com/peteraba/d5/lib/german"
)

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func parseDictionary(dictionary [][6]string, user string) ([]germanLib.Word, []string) {
	var (
		words       = []germanLib.Word{}
		parseErrors = []string{}
	)

	for _, word := range dictionary {
		var (
			w        germanLib.Word
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
			w = germanLib.NewAdjective(german, english, third, user, learned, score)
			break
		case "noun":
			if germanLib.NounRegexp.MatchString(german) {
				w = germanLib.NewNoun(german, english, third, user, learned, score)
			}
			break
		case "verb":
			if germanLib.VerbRegexp.MatchString(german) {
				w = germanLib.NewVerb(german, english, third, user, learned, score)
			}
			break
		default:
			w = germanLib.NewAny(german, english, third, category, user, learned, score, true)
		}

		if w == nil {
			parseErrors = append(parseErrors, german)

			w = germanLib.NewAny(german, english, third, category, user, learned, score, false)
		}

		words = append(words, w)
	}

	return words, parseErrors
}

func main() {
	var (
		user       = ""
		logErrors  = true
		dictionary = [][6]string{}
	)

	if len(os.Args) > 1 {
		user = os.Args[1]
	}
	if len(os.Args) > 2 {
		logErrors, _ = strconv.ParseBool(os.Args[2])
	}

	input, err := readStdInput()
	if err != nil && logErrors {
		log.Println(err)
	}

	json.Unmarshal(input, &dictionary)

	words, parseErrors := parseDictionary(dictionary, user)

	if logErrors && len(parseErrors) > 0 {
		for _, word := range parseErrors {
			fmt.Printf("Failed: %v\n", word)
		}
	}

	b, err := json.Marshal(words)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
