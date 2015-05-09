package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	germanEntity "github.com/peteraba/d5/lib/german/entity"
)

func readStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func parseDictionary(dictionary [][8]string, user string) ([]germanEntity.Word, []string) {
	var (
		words       = []germanEntity.Word{}
		parseErrors = []string{}
	)

	for _, word := range dictionary {
		var (
			w                  germanEntity.Word
			articleOrAuxiliary = word[0]
			german             = word[1]
			english            = word[2]
			third              = word[3]
			category           = word[4]
			learned            = word[5]
			score              = word[6]
			tags               = word[7]
		)

		if english == "" {
			continue
		}

		switch category {
		case "adj":
			w = germanEntity.NewAdjective(german, english, third, user, learned, score, tags)
			break
		case "noun":
			if germanEntity.NounRegexp.MatchString(german) {
				w = germanEntity.NewNoun(articleOrAuxiliary, german, english, third, user, learned, score, tags)
			}
			break
		case "verb":
			if germanEntity.VerbRegexp.MatchString(german) {
				w = germanEntity.NewVerb(articleOrAuxiliary, german, english, third, user, learned, score, tags)
			}
			break
		default:
			w = germanEntity.NewAny(german, english, third, category, user, learned, score, tags, []string{})
		}

		if w == nil {

			w = germanEntity.NewAny(german, english, third, category, user, learned, score, tags, []string{"Parsing failed."})
			if w == nil {
				parseErrors = append(parseErrors, german)
				continue
			} else {
				parseErrors = append(parseErrors, german+"!!!!!")
			}
		}

		words = append(words, w)
	}

	return words, parseErrors
}

func parseFlags() (string, bool) {
	user := flag.String("user", "", "User to whom the data belongs")

	debug := flag.Bool("debug", false, "Log errors, halt output")

	flag.Parse()

	return *user, *debug
}

func main() {
	var (
		dictionary = [][8]string{}
	)

	user, debug := parseFlags()

	if user == "" {
		return
	}

	input, err := readStdInput()
	if err != nil && debug {
		log.Println(err)
	}

	json.Unmarshal(input, &dictionary)

	words, parseErrors := parseDictionary(dictionary, user)

	if debug && len(parseErrors) > 0 {
		for _, word := range parseErrors {
			log.Printf("Failed: %v\n", word)
		}
	}

	b, err := json.Marshal(words)
	if err != nil {
		log.Fatalln(err)
	}

	if !debug {
		fmt.Println(string(b))
	} else {
		log.Println("Parsing is done.")
	}
}
