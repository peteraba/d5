package entity

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/peteraba/d5/lib/util"
)

const learnedForm = "2006-01-02"

const (
	alternativeSeparator = "/"
	conjugationSeparator = ","
	meaningSeparator     = ";"
	synonimSeparator     = ","
	tagSeparator         = ","
	defaultWhitespace    = "\t\n\f\r "
	wordSeparator        = " "
)

var (
	// Meaning:
	// ^                                -- match beginning of string
	//  ([^(]*)                         -- match any string that is not a open parantheses character
	//          ?                       -- match optional space character
	//           (                      -- start of parantheses matching
	//            [(]                     -- match open parantheses character
	//               ([^)]*)              -- match parantheses content
	//                      [)]           -- match close parantheses character
	//                         )?         -- end of parantheses matching
	//                           *      -- match optional spaces
	//                            $     -- match end of string
	MeaningRegexp = regexp.MustCompile("^([^(]*) ?([(]([^)]*)[)])? *$")

	// English Word:
	// ^                       -- match beginning of string
	//  [a-zA-Z,.() ]*         -- English words can only contain letters, dots, parantheses and spaces
	//                $        -- match end of string
	EnglishRegexp = regexp.MustCompile("^[a-zA-Z,.() ]*$")

	// German Word:
	// ^                           -- match beginning of string
	//  [a-zA-ZäÄöÖüÜß,.() ]*      -- German words can only contain German letters, dots, parantheses and spaces
	//                       $     -- match end of string
	GermanRegexp = regexp.MustCompile("^[a-zA-ZäÄöÖüÜß,.() ]*$")
)

type Word interface {
	GetGerman() string
	GetEnglish() []Meaning
	GetThird() []Meaning
	GetCategory() string
	GetUser() string
	GetScore() int
	GetLearned() time.Time
	GetErrors() []string
}

type Meaning struct {
	Main        string `bson:"main" json:"main"`
	Parantheses string `bson:"parantheses" json:"parantheses"`
}

func NewMeanings(allMeanings string) []Meaning {
	meanings := []Meaning{}

	for _, word := range util.TrimSplit(allMeanings, meaningSeparator) {
		matches := MeaningRegexp.FindStringSubmatch(word)

		m := strings.Trim(matches[1], defaultWhitespace)
		p := strings.Trim(matches[3], defaultWhitespace)

		meanings = append(meanings, Meaning{m, p})
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
	Tags     []string  `bson:"tags" json:"tags"`
	Errors   []string  `bson:"errors" json:"errors"`
}

func NewDefaultWord(german, english, third, category, user, learned, score, tags string, errors []string) DefaultWord {
	englishMeanings, thirdMeanings := NewMeanings(english), NewMeanings(third)

	scoreParsed, err := strconv.ParseInt(score, 0, 0)
	if err != nil || scoreParsed < 1 || scoreParsed > 10 {
		scoreParsed = 5
	}

	learnedParsed := util.ParseTimeNow(learnedForm, learned)

	return DefaultWord{
		german,
		englishMeanings,
		thirdMeanings,
		category,
		user,
		learnedParsed,
		int(scoreParsed),
		util.TrimSplit(tags, tagSeparator),
		errors,
	}
}

func (w DefaultWord) GetGerman() string {
	return w.German
}

func (w DefaultWord) GetEnglish() []Meaning {
	return w.English
}

func (w DefaultWord) GetThird() []Meaning {
	return w.Third
}

func (w DefaultWord) GetCategory() string {
	return w.Category
}

func (w DefaultWord) GetScore() int {
	return w.Score
}

func (w DefaultWord) GetUser() string {
	return w.User
}

func (w DefaultWord) GetLearned() time.Time {
	return w.Learned
}

func (w DefaultWord) GetErrors() []string {
	return w.Errors
}

type Any struct {
	DefaultWord `bson:"word" json:"word"`
}

func NewAny(german, english, third, category, user, learned, score, tags string, errors []string) *Any {
	d := NewDefaultWord(german, english, third, category, user, learned, score, tags, errors)

	return &Any{d}
}
