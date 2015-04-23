package main

import (
	"log"
	"strings"

	"gopkg.in/mgo.v2"
)

type Meaning struct {
	Main        string `bson:"main" json:"main"`
	Paranthases string `bson:"paranthases" json:"paranthases"`
}

type Word struct {
	German   string    `bson:"german" json:"german"`
	English  []Meaning `bson:"english" json:"english"`
	Third    []Meaning `bson:"third" json:"third"`
	Category string    `bson:"category" json:"category"`
}

type Verb struct {
	Word           `bson:"word" json:"word"`
	Auxiliary      []string `bson:"auxiliary" json:"auxiliary"`
	PastParticiple string   `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      string   `bson:"preterite" json:"preterite"`
	Du             string   `bson:"du" json:"du"`
	Er             string   `bson:"er" json:"er"`
}

type Noun struct {
	Word   `bson:"word" json:"word"`
	Plural []string `bson:"plural" json:"plural"`
}

type Adjective struct {
	Word        `bson:"word" json:"word"`
	Comparative []string `bson:"comparative" json:"comparative"`
	Superlative []string `bson:"superlative" json:"superlative"`
}

func NewMeanings(allMeanings string) []Meaning {
	meanings := []Meaning{}

	for _, word := range strings.Split(allMeanings, ";") {
		meanings = append(meanings, Meaning{word, ""})
	}

	return meanings
}

func NewWord(german, english, third, category string) Word {
	return Word{german, NewMeanings(english), NewMeanings(third), category}
}

func NewVerb(german, english, third string) Verb {
	auxiliary := []string{}
	pastParticiple := ""
	preterite := ""
	du := ""
	er := ""

	return Verb{
		NewWord(german, english, third, "verb"),
		auxiliary,
		pastParticiple,
		preterite,
		du,
		er,
	}
}

func NewNoun(german, english, third string) Noun {
	return Noun{
		NewWord(german, english, third, "noun"),
		[]string{},
	}
}

func NewAdjective(german, english, third string) Adjective {
	adjectiveParts := strings.Split(german, ",")

	comparative := []string{}
	superlative := []string{}

	if len(adjectiveParts) > 1 {
		comparative = strings.Split(adjectiveParts[1], "/")
	}
	if len(adjectiveParts) > 2 {
		superlative = strings.Split(adjectiveParts[2], "/")
	}

	return Adjective{
		NewWord(german, english, third, "adjective"),
		comparative,
		superlative,
	}
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c5 := session.DB("test").C("meaning")
	err = c5.Insert(Meaning{"abc", "d"}, Meaning{"qwe", "r"})
	if err != nil {
		log.Fatal(err)
	}

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
	err = c4.Insert(a1, a2, a3)
	if err != nil {
		log.Fatal(err)
	}
}
