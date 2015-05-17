package german

import (
	"encoding/json"

	"github.com/peteraba/d5/lib/german/entity"
)

type Superword struct {
	entity.DefaultWord `bson:"word" json:"word"`
	Auxiliary          []entity.Auxiliary `bson:"auxiliary" json:"auxiliary,omitempty"`
	Prefix             entity.Prefix      `bson:"prefix" json:"prefix,omitempty"`
	Noun               string             `bson:"noun" json:"noun,omitempty"`
	Adjective          string             `bson:"adjective" json:"adjective,omitempty"`
	PastParticiple     []string           `bson:"pastParticiple" json:"pastParticiple,omitempty"`
	Preterite          []string           `bson:"preterite" json:"preterite,omitempty"`
	S1                 []string           `bson:"s1" json:"s1,omitempty"`
	S2                 []string           `bson:"s2" json:"s2,omitempty"`
	S3                 []string           `bson:"s3" json:"s3,omitempty"`
	P1                 []string           `bson:"p1" json:"p1,omitempty"`
	P2                 []string           `bson:"p2" json:"p2,omitempty"`
	P3                 []string           `bson:"p3" json:"p3,omitempty"`
	Reflexive          entity.Reflexive   `bson:"reflexive" json:"reflexive,omitempty"`
	Arguments          []entity.Argument  `bson:"arguments" json:"arguments,omitempty"`
	Articles           []entity.Article   `bson:"article" json:"article,omitempty"`
	Plural             []string           `bson:"plural" json:"plural,omitempty"`
	Genitive           []string           `bson:"genitive" json:"genitive,omitempty"`
	IsPluralOnly       bool               `bson:"plural_only" json:"plural_only,omitempty"`
	Comparative        []string           `bson:"comparative" json:"comparative,omitempty"`
	Superlative        []string           `bson:"superlative" json:"superlative,omitempty"`
}

type Dictionary struct {
	Nouns      []entity.Noun           `bson:"nouns" json:"nouns,omitempty"`
	Verbs      []entity.Verb           `bson:"verbs" json:"verbs,omitempty"`
	Adjectives []entity.Adjective      `bson:"adjectives" json:"adjectives,omitempty"`
	Words      map[string][]entity.Any `bson:"words" json:"words,omitempty"`
}

func NewDictionary() Dictionary {
	var dictionary = Dictionary{}

	dictionary.Words = map[string][]entity.Any{}

	return dictionary
}

func ParseWords(input []byte) ([]entity.Word, error) {
	var (
		superwords = []Superword{}
		words      = []entity.Word{}
		err        error
	)

	if err = json.Unmarshal(input, &superwords); err != nil {
		return words, err
	}

	return SuperwordsToWords(superwords)
}

func SuperwordsToWords(superwords []Superword) ([]entity.Word, error) {
	var (
		words = []entity.Word{}
		word  entity.Word
	)

	for _, superword := range superwords {
		switch superword.Category {
		case "verb":
			word = &entity.Verb{
				superword.DefaultWord,
				superword.Auxiliary,
				superword.Prefix,
				superword.Noun,
				superword.Adjective,
				superword.PastParticiple,
				superword.Preterite,
				superword.S1,
				superword.S2,
				superword.S3,
				superword.P1,
				superword.P2,
				superword.P3,
				superword.Reflexive,
				superword.Arguments,
			}
			break
		case "noun":
			word = &entity.Noun{
				superword.DefaultWord,
				superword.Articles,
				superword.Plural,
				superword.Genitive,
				superword.IsPluralOnly,
			}
			break
		case "adj":
			word = &entity.Adjective{
				superword.DefaultWord,
				superword.Comparative,
				superword.Superlative,
			}
			break
		default:
			word = &entity.Any{
				superword.DefaultWord,
			}
			break
		}

		words = append(words, word)
	}

	return words, nil
}

func SuperwordsToDictionary(superwords []Superword) Dictionary {
	var (
		dictionary = NewDictionary()
	)

	for _, superword := range superwords {
		cat := superword.Category

		switch cat {
		case "verb":
			dictionary.Verbs = append(
				dictionary.Verbs,
				entity.Verb{
					superword.DefaultWord,
					superword.Auxiliary,
					superword.Prefix,
					superword.Noun,
					superword.Adjective,
					superword.PastParticiple,
					superword.Preterite,
					superword.S1,
					superword.S2,
					superword.S3,
					superword.P1,
					superword.P2,
					superword.P3,
					superword.Reflexive,
					superword.Arguments,
				},
			)
			break
		case "noun":
			dictionary.Nouns = append(
				dictionary.Nouns,
				entity.Noun{
					superword.DefaultWord,
					superword.Articles,
					superword.Plural,
					superword.Genitive,
					superword.IsPluralOnly,
				},
			)
			break
		case "adj":
			dictionary.Adjectives = append(
				dictionary.Adjectives,
				entity.Adjective{
					superword.DefaultWord,
					superword.Comparative,
					superword.Superlative,
				},
			)
			break
		default:
			if _, ok := dictionary.Words[cat]; !ok {
				dictionary.Words[cat] = []entity.Any{}
			}

			dictionary.Words[cat] = append(
				dictionary.Words[cat],
				entity.Any{
					superword.DefaultWord,
				},
			)
			break
		}
	}

	return dictionary
}

func (d *Dictionary) GetCount() int {
	var count = 0

	count = len(d.Verbs) + len(d.Nouns) + len(d.Adjectives)

	for _, words := range d.Words {
		count += len(words)
	}

	return count
}
