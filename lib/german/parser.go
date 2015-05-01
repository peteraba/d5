package german

import "encoding/json"

type superword struct {
	DefaultWord    `bson:"word" json:"word"`
	Auxiliary      []string  `bson:"auxiliary" json:"auxiliary"`
	PastParticiple []string  `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      []string  `bson:"preterite" json:"preterite"`
	Ich            []string  `bson:"ich" json:"ich"`
	Du             []string  `bson:"du" json:"du"`
	Er             []string  `bson:"er" json:"er"`
	Wir            []string  `bson:"wir" json:"wir"`
	Ihr            []string  `bson:"ihr" json:"ihr"`
	Sie            []string  `bson:"sie" json:"sie"`
	Sich           Sich      `bson:"sich" json:"sich"`
	Arguments      []string  `bson:"arguments" json:"arguments"`
	Articles       []Article `bson:"article" json:"article"`
	Plural         []string  `bson:"plural" json:"plural"`
	Comparative    []string  `bson:"comparative" json:"comparative"`
	Superlative    []string  `bson:"superlative" json:"superlative"`
}

func ParseWords(input []byte) ([]Word, error) {
	var (
		superwords = []superword{}
		words      = []Word{}
		word       Word
		err        error
	)

	if err = json.Unmarshal(input, &superwords); err != nil {
		return words, err
	}

	for _, superword := range superwords {
		switch superword.Category {
		case "verb":
			word = &Verb{
				superword.DefaultWord,
				superword.Auxiliary,
				superword.PastParticiple,
				superword.Preterite,
				superword.Ich,
				superword.Du,
				superword.Er,
				superword.Wir,
				superword.Ihr,
				superword.Sie,
				superword.Sich,
				superword.Arguments,
			}
			break
		case "noun":
			word = &Noun{
				superword.DefaultWord,
				superword.Articles,
				superword.Plural,
			}
			break
		case "adjective":
			word = &Adjective{
				superword.DefaultWord,
				superword.Comparative,
				superword.Superlative,
			}
			break
		default:
			word = &Any{
				superword.DefaultWord,
			}
			break
		}

		words = append(words, word)
	}

	return words, nil
}
