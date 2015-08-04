package german

import (
	"encoding/json"
	"fmt"

	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
)

type Superword struct {
	entity.DefaultWord `bson:"word" json:"word"`
	Id                 bson.ObjectId      `bson:"_id,omitempty" json:"_id,omitempty"`
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

func (s Superword) GetId() string {
	return fmt.Sprintf("%x", string(s.Id))
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

	return SuperwordsToWords(superwords), nil
}

func SuperwordsToWords(superwords []Superword) []entity.Word {
	var (
		words = []entity.Word{}
		word  entity.Word
	)

	for _, superword := range superwords {
		switch superword.Category {
		case "verb":
			verb := superwordToVerb(superword)

			word = &verb

			break
		case "noun":
			noun := superwordToNoun(superword)

			word = &noun

			break
		case "adj":
			adjective := superwordToAdjective(superword)

			word = &adjective

			break
		default:
			any := superwordToAny(superword)

			word = &any

			break
		}

		words = append(words, word)
	}

	return words
}

func SuperwordsToDictionary(superwords []Superword) Dictionary {
	var (
		dictionary = NewDictionary()
	)

	for _, superword := range superwords {
		cat := superword.Category

		switch cat {
		case "verb":
			verb := superwordToVerb(superword)

			dictionary.Verbs = append(dictionary.Verbs, verb)

			break
		case "noun":
			noun := superwordToNoun(superword)

			dictionary.Nouns = append(dictionary.Nouns, noun)

			break
		case "adj":
			adjective := superwordToAdjective(superword)

			dictionary.Adjectives = append(dictionary.Adjectives, adjective)

			break
		default:
			if _, ok := dictionary.Words[cat]; !ok {
				dictionary.Words[cat] = []entity.Any{}
			}

			any := superwordToAny(superword)

			dictionary.Words[cat] = append(dictionary.Words[cat], any)

			break
		}
	}

	return dictionary
}

func superwordToNoun(superword Superword) entity.Noun {
	noun := entity.Noun{}

	noun.DefaultWord.German = superword.DefaultWord.German
	noun.DefaultWord.English = superword.DefaultWord.English
	noun.DefaultWord.Third = superword.DefaultWord.Third
	noun.DefaultWord.Category = superword.DefaultWord.Category
	noun.DefaultWord.User = superword.DefaultWord.User
	noun.DefaultWord.Learned = superword.DefaultWord.Learned
	noun.DefaultWord.Score = superword.DefaultWord.Score
	noun.DefaultWord.Tags = superword.DefaultWord.Tags
	noun.DefaultWord.Errors = superword.DefaultWord.Errors
	noun.DefaultWord.Scores = superword.DefaultWord.Scores

	noun.SetId(superword.GetId())

	noun.Articles = superword.Articles
	noun.Plural = superword.Plural
	noun.Genitive = superword.Genitive
	noun.IsPluralOnly = superword.IsPluralOnly

	return noun
}

func superwordToVerb(superword Superword) entity.Verb {
	verb := entity.Verb{}

	verb.DefaultWord.German = superword.DefaultWord.German
	verb.DefaultWord.English = superword.DefaultWord.English
	verb.DefaultWord.Third = superword.DefaultWord.Third
	verb.DefaultWord.Category = superword.DefaultWord.Category
	verb.DefaultWord.User = superword.DefaultWord.User
	verb.DefaultWord.Learned = superword.DefaultWord.Learned
	verb.DefaultWord.Score = superword.DefaultWord.Score
	verb.DefaultWord.Tags = superword.DefaultWord.Tags
	verb.DefaultWord.Errors = superword.DefaultWord.Errors
	verb.DefaultWord.Scores = superword.DefaultWord.Scores

	verb.SetId(superword.GetId())

	verb.Auxiliary = superword.Auxiliary
	verb.Prefix = superword.Prefix
	verb.Noun = superword.Noun
	verb.Adjective = superword.Adjective
	verb.PastParticiple = superword.PastParticiple
	verb.Preterite = superword.Preterite
	verb.S1 = superword.S1
	verb.S2 = superword.S2
	verb.S3 = superword.S3
	verb.P1 = superword.P1
	verb.P2 = superword.P2
	verb.P3 = superword.P3
	verb.Reflexive = superword.Reflexive
	verb.Arguments = superword.Arguments

	return verb
}

func superwordToAdjective(superword Superword) entity.Adjective {
	adjective := entity.Adjective{}

	adjective.DefaultWord.German = superword.DefaultWord.German
	adjective.DefaultWord.English = superword.DefaultWord.English
	adjective.DefaultWord.Third = superword.DefaultWord.Third
	adjective.DefaultWord.Category = superword.DefaultWord.Category
	adjective.DefaultWord.User = superword.DefaultWord.User
	adjective.DefaultWord.Learned = superword.DefaultWord.Learned
	adjective.DefaultWord.Score = superword.DefaultWord.Score
	adjective.DefaultWord.Tags = superword.DefaultWord.Tags
	adjective.DefaultWord.Errors = superword.DefaultWord.Errors
	adjective.DefaultWord.Scores = superword.DefaultWord.Scores

	adjective.SetId(superword.GetId())

	adjective.Comparative = superword.Comparative
	adjective.Superlative = superword.Superlative

	return adjective
}

func superwordToAny(superword Superword) entity.Any {
	any := entity.Any{}

	any.DefaultWord.German = superword.DefaultWord.German
	any.DefaultWord.English = superword.DefaultWord.English
	any.DefaultWord.Third = superword.DefaultWord.Third
	any.DefaultWord.Category = superword.DefaultWord.Category
	any.DefaultWord.User = superword.DefaultWord.User
	any.DefaultWord.Learned = superword.DefaultWord.Learned
	any.DefaultWord.Score = superword.DefaultWord.Score
	any.DefaultWord.Tags = superword.DefaultWord.Tags
	any.DefaultWord.Errors = superword.DefaultWord.Errors
	any.DefaultWord.Scores = superword.DefaultWord.Scores

	any.SetId(superword.GetId())

	return any
}

func (d *Dictionary) GetCount() int {
	var count = 0

	count = len(d.Verbs) + len(d.Nouns) + len(d.Adjectives)

	for _, words := range d.Words {
		count += len(words)
	}

	return count
}
