package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/peteraba/d5/shared"
)

func main() {
	var (
		user       = os.Args[1]
		input      = []byte{}
		dictionary = [][6]string{}
		words      = []shared.Word{}
	)

	fmt.Scan(&input)

	json.Unmarshal(input, &dictionary)

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
			fmt.Printf("Failed: %v\n", german)
			w = shared.NewWord(german, english, third, category, user, learned, score, false)
		}

		words = append(words, w)
	}
}
