package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Sich int

const (
	Without Sich = iota
	Acusative
	Dative
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
	PastParticiple string   `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      string   `bson:"preterite" json:"preterite"`
	Du             string   `bson:"du" json:"du"`
	Er             string   `bson:"er" json:"er"`
	Sich           Sich     `bson:"sich" json:"sich"`
	Arguments      []string `bson:"arguments" json:"arguments"`
}

func NewVerb(german, english, third, user string, multiplier int, verbRegexp *regexp.Regexp) *Verb {
	pastParticiple := ""
	preterite := ""
	du := ""
	er := ""
	sich := Without
	arguments := []string{}

	matches := verbRegexp.FindStringSubmatch(german)

	auxiliary := strings.Split(matches[1], "/")

	main := strings.Split(matches[2], ",")

	german = main[0]

	if len(main) == 2 || len(main) == 4 || len(main) > 5 {
		return nil
	}
	if len(main) > 2 {
		pastParticiple = main[1]
		preterite = main[2]
	}
	if len(main) > 3 {
		du = main[3]
		er = main[4]
	}

	if matches[3] != "" {
		arguments = strings.Split(matches[3], "+")

		if strings.Contains(arguments[0], "sich (A)") {
			arguments = arguments[1:]
			sich = Acusative
		}

		if strings.Contains(arguments[0], "sich (D)") {
			arguments = arguments[1:]
			sich = Dative
		}

		if strings.Contains(arguments[0], "sich") {
			return nil
		}
	}

	return &Verb{
		NewDefaultWord(german, english, third, "verb", user, multiplier),
		auxiliary,
		pastParticiple,
		preterite,
		du,
		er,
		sich,
		arguments,
	}
}

type Noun struct {
	DefaultWord `bson:"word" json:"word"`
	Plural      []string `bson:"plural" json:"plural"`
}

func NewNoun(german, english, third, user string, multiplier int, nounRegexp *regexp.Regexp) *Noun {
	matches := nounRegexp.FindAllStringSubmatch(german, -1)

	fmt.Printf("Matches: %q\n", matches)

	return &Noun{
		NewDefaultWord(german, english, third, "noun", user, multiplier),
		[]string{},
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

	verbRegexp := regexp.MustCompile("^([sh/]+) ([^+]*)([+](.*))$")

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
