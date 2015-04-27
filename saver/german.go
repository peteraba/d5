package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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
	German     string    `bson:"german" json:"german"`
	English    []Meaning `bson:"english" json:"english"`
	Third      []Meaning `bson:"third" json:"third"`
	Category   string    `bson:"category" json:"category"`
	User       string    `bson:"user" json:"user"`
	Multiplier int       `bson:"multiplier" json:"multiplier"`
	Ok         bool      `bson:"ok", json:"ok"`
}

func NewDefaultWord(german, english, third, category, user string, multiplier int) DefaultWord {
	return DefaultWord{german, NewMeanings(english), NewMeanings(third), category, user, multiplier, true}
}

func NewWord(german, english, third, category, user string, multiplier int, ok bool) *DefaultWord {
	return &DefaultWord{german, NewMeanings(english), NewMeanings(third), category, user, multiplier, ok}
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

func NewVerb(german, english, third, user string, multiplier int, verbRegexp *regexp.Regexp) *Verb {
	pastParticiple := ""
	preterite := ""
	ich := ""
	du := ""
	er := ""
	wir := ""
	ihr := ""
	sie := ""

	matches := verbRegexp.FindStringSubmatch(german)

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
		NewDefaultWord(german, english, third, "verb", user, multiplier),
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

func NewNoun(german, english, third, user string, multiplier int, nounRegexp *regexp.Regexp) *Noun {
	matches := nounRegexp.FindStringSubmatch(german)

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
		NewDefaultWord(german, english, third, "noun", user, multiplier),
		articles,
		strings.Split(matches[3], "/"),
	}
}

type Adjective struct {
	DefaultWord `bson:"word" json:"word"`
	Comparative []string `bson:"comparative" json:"comparative"`
	Superlative []string `bson:"superlative" json:"superlative"`
}

func NewAdjective(german, english, third, user string, multiplier int) *Adjective {
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
		NewDefaultWord(german, english, third, "adjective", user, multiplier),
		comparative,
		superlative,
	}
}

func main() {
	/*
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)
	*/
	nounRegexp := regexp.MustCompile("^([ers/]+) ([^,]+),(.*)")

	verbRegexp := regexp.MustCompile("^([sh/]+) ([^+]*)([+](.*))?$")

	file, e := ioutil.ReadFile("./test.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//m := new(Dispatch)
	//var m interface{}
	dictionary := [][4]string{}
	json.Unmarshal(file, &dictionary)
	//fmt.Printf("Results: %v\n", dictionary)

	words := []Word{}

	multiplier := 5

	user := "peteraba"

	for _, word := range dictionary {
		var w Word

		if word[1] == "" {
			continue
		}

		switch word[3] {
		case "adj":
			w = NewAdjective(word[0], word[1], word[2], user, multiplier)
			break
		case "noun":
			if nounRegexp.MatchString(word[0]) {
				w = NewNoun(word[0], word[1], word[2], user, multiplier, nounRegexp)
			}
			break
		case "verb":
			if verbRegexp.MatchString(word[0]) {
				w = NewVerb(word[0], word[1], word[2], user, multiplier, verbRegexp)
			}
			break
		default:
			w = NewWord(word[0], word[1], word[2], word[3], user, multiplier, true)
		}

		if w == nil {
			fmt.Printf("Failed: %v\n", word[0])
			w = NewWord(word[0], word[1], word[2], word[3], user, multiplier, false)
		}

		words = append(words, w)
	}

	/*for _, w := range words {
		fmt.Printf("Results: %v\n", w)
	}
	/*
		w1 := NewWord("h runternehmen, runternahmen, runtergenommen, runternimmst, runternimmt", "to take down", "levenni vmit (föntről)", "verb")
		w2 := NewWord("schlecht sein + für (A)", "to do wrong to sth", "rosszat tenni vminek", "verb")
		w3 := NewWord("h scheren, schoren/scherten, geschoren/geschert", "to shear (sheep), to cut, to trim (hair)", "nyírni, lenyírni", "verb")
		c1 := session.DB("test").C("word")
		err = c1.Insert(&w1, &w2, &w3, w1, w2, w3)
		if err != nil {
			log.Fatal(err)
		}

		v1 := NewVerb("h runternehmen, runternahmen, runtergenommen, runternimmst, runternimmt", "to take down", "levenni vmit (föntről)")
		v2 := NewVerb("schlecht sein + für (A)", "to do wrong to sth", "rosszat tenni vminek")
		v3 := NewVerb("h scheren, schoren/scherten, geschoren/geschert", "to shear (sheep), to cut, to trim (hair)", "nyírni, lenyírni")
		c2 := session.DB("test").C("verb")
		err = c2.Insert(&v1, &v2, &v3)
		if err != nil {
			log.Fatal(err)
		}

		n1 := NewNoun("r Arbeitsplatz,⍨e", "work place (both abstract or concrete)", "munkahely (absztrakt vagy konkrét)")
		n2 := NewNoun("r/s Hot Dog,~s", "hotdog", "hotdog")
		n3 := NewNoun("s Jurastudium, Jurastudien", "law studies", "jogi tanulmány")
		c3 := session.DB("test").C("noun")
		err = c3.Insert(&n1, &n2, &n3)
		if err != nil {
			log.Fatal(err)
		}

		a1 := NewAdjective("weich,~er,~esten", "soft", "puha")
		a2 := NewAdjective("ständig,-", "persistent; permanent", "állandó; állandóan, folyton")
		a3 := NewAdjective("schmal,~er/⍨er,~sten/⍨sten", "narrow", "keskeny, szűk")
		c4 := session.DB("test").C("adjective")
		err = c4.Insert(a1, a2, a3	if err != nil {
			log.Fatal(err)
		}*/
}

func log(err error) {
	fmt.Println(err)
}
