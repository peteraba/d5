package shared

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Sich string

const (
	Without   Sich = ""
	Acusative      = "a"
	Dative         = "d"
)

type Article string

const (
	Der Article = "r"
	Die         = "e"
	Das         = "s"
)

const learnedForm = "2006-01-02"

var (
	NounRegexp = regexp.MustCompile("^([ers/]+) ([^,]+),(.*)$")
	VerbRegexp = regexp.MustCompile("^([sh/]+) ([^+]*)([+](.*))?$")
)

type Word interface {
	GetGerman() string
	GetEnglish() []Meaning
	GetThird() []Meaning
	GetCategory() string
}

type Meaning struct {
	Main        string `bson:"main" json:"main"`
	Paranthases string `bson:"paranthases" json:"paranthases"`
}

func NewMeanings(allMeanings string) []Meaning {
	meanings := []Meaning{}

	for _, word := range strings.Split(allMeanings, ";") {
		meanings = append(meanings, Meaning{word, ""})
	}

	return meanings
}

type DefaultWord struct {
	German   string    `bson:"german" json:"german"`
	English  []Meaning `bson:"english" json:"english"`
	Third    []Meaning `bson:"third" json:"third"`
	Category string    `bson:"category" json:"category"`
	User     string    `bson:"user" json:"user"`
	Learned  time.Time `bson:"learned" json:"learned"`
	Score    int       `bson:"score" json:"score"`
	Ok       bool      `bson:"ok", json:"ok"`
}

func NewDefaultWord(german, english, third, category, user, learned, score string) DefaultWord {
	englishMeanings, thirdMeanings := NewMeanings(english), NewMeanings(third)

	scoreParsed, err := strconv.ParseInt(score, 0, 0)
	if err != nil || scoreParsed < 0 || scoreParsed > 10 {
		scoreParsed = 5
	}

	learnedParsed, err := time.Parse(learnedForm, learned)
	if err != nil {
		learnedParsed = time.Now()
	}

	return DefaultWord{german, englishMeanings, thirdMeanings, category, user, learnedParsed, int(scoreParsed), true}
}

func NewWord(german, english, third, category, user, learned, score string, ok bool) *DefaultWord {
	d := NewDefaultWord(german, english, third, category, user, learned, score)

	d.Ok = ok

	return &d
}

func (w *DefaultWord) GetGerman() string {
	return w.German
}

func (w *DefaultWord) GetEnglish() []Meaning {
	return w.English
}

func (w *DefaultWord) GetThird() []Meaning {
	return w.Third
}

func (w *DefaultWord) GetCategory() string {
	return w.Category
}

type Verb struct {
	DefaultWord    `bson:"word" json:"word"`
	Auxiliary      []string `bson:"auxiliary" json:"auxiliary"`
	PastParticiple []string `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      []string `bson:"preterite" json:"preterite"`
	Ich            []string `bson:"ich" json:"ich"`
	Du             []string `bson:"du" json:"du"`
	Er             []string `bson:"er" json:"er"`
	Wir            []string `bson:"wir" json:"wir"`
	Ihr            []string `bson:"ihr" json:"ihr"`
	Sie            []string `bson:"sie" json:"sie"`
	Sich           Sich     `bson:"sich" json:"sich"`
	Arguments      []string `bson:"arguments" json:"arguments"`
}

type Superverb struct {
	Verb `bson:"verb" json:"verb"`
}

func NewVerb(german, english, third, user, learned, score string) *Verb {
	pastParticiple := ""
	preterite := ""
	ich := ""
	du := ""
	er := ""
	wir := ""
	ihr := ""
	sie := ""

	matches := VerbRegexp.FindStringSubmatch(german)

	main := strings.Split(matches[2], ",")
	switch len(main) {
	case 1:
		german = main[0]
		break
	case 3:
		german, pastParticiple, preterite = main[0], main[1], main[2]
		break
	case 5:
		german, pastParticiple, preterite, du, er = main[0], main[1], main[2], main[3], main[4]
		break
	case 9:
		german, ich, du, er, wir, ihr, sie, pastParticiple, preterite = main[0], main[1], main[2], main[3], main[4], main[5], main[6], main[7], main[8]
		break
	default:
		return nil
	}

	sich, arguments, err := parseArguments(matches[3])
	if err != nil {
		return nil
	}

	return &Verb{
		NewDefaultWord(german, english, third, "verb", user, learned, score),
		strings.Split(matches[1], "/"),
		strings.Split(pastParticiple, "/"),
		strings.Split(preterite, "/"),
		strings.Split(ich, "/"),
		strings.Split(du, "/"),
		strings.Split(er, "/"),
		strings.Split(wir, "/"),
		strings.Split(ihr, "/"),
		strings.Split(sie, "/"),
		sich,
		arguments,
	}
}

func parseArguments(rawArguments string) (Sich, []string, error) {
	if rawArguments == "" {
		return Without, []string{}, nil
	}
	arguments := strings.Split(rawArguments, "+")

	if strings.Contains(arguments[0], "sich (A)") {
		return Acusative, arguments[1:], nil
	}

	if strings.Contains(arguments[0], "sich (D)") {
		return Dative, arguments[1:], nil
	}

	if strings.Contains(arguments[0], "sich") {
		return Without, []string{}, errors.New("Sich definition is invalid")
	}

	return Without, []string{}, nil
}

type Noun struct {
	DefaultWord `bson:"word" json:"word"`
	Articles    []Article `bson:"article" json:"article"`
	Plural      []string  `bson:"plural" json:"plural"`
}

func NewNoun(german, english, third, user, learned, score string) *Noun {
	matches := NounRegexp.FindStringSubmatch(german)

	articles := []Article{}
	for _, article := range strings.Split(matches[1], "/") {
		switch article {
		case "r":
			articles = append(articles, Der)
			break
		case "e":
			articles = append(articles, Die)
			break
		case "s":
			articles = append(articles, Das)
			break
		default:
			return nil
		}
	}

	return &Noun{
		NewDefaultWord(german, english, third, "noun", user, learned, score),
		articles,
		strings.Split(matches[3], "/"),
	}
}

type Adjective struct {
	DefaultWord `bson:"word" json:"word"`
	Comparative []string `bson:"comparative" json:"comparative"`
	Superlative []string `bson:"superlative" json:"superlative"`
}

func NewAdjective(german, english, third, user, learned, score string) *Adjective {
	adjectiveParts := strings.Split(german, ",")

	comparative := []string{}
	superlative := []string{}

	if len(adjectiveParts) > 1 {
		comparative = strings.Split(adjectiveParts[1], "/")
	}
	if len(adjectiveParts) > 2 {
		superlative = strings.Split(adjectiveParts[2], "/")
	}

	return &Adjective{
		NewDefaultWord(german, english, third, "adjective", user, learned, score),
		comparative,
		superlative,
	}
}
