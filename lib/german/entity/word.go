package entity

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/peteraba/d5/lib/general"
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

	// German Word:
	// ^                           -- match beginning of string
	//  [a-zA-ZäÄöÖüÜß,.() ]*      -- German words can only contain German letters, dots, parantheses and spaces
	//                       $     -- match end of string
	GermanRegexp = regexp.MustCompile("^[a-zA-ZäÄöÖüÜß,.() ]*$")
)

type Word interface {
	GetId() string
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
	Main        string `bson:"main" json:"main,omitempty"`
	Parantheses string `bson:"parantheses" json:"parantheses,omitempty"`
}

func NewMeanings(allMeanings string, errors []string) ([]Meaning, []string) {
	meanings := []Meaning{}

	for _, word := range util.TrimSplit(allMeanings, meaningSeparator) {
		matches := MeaningRegexp.FindStringSubmatch(word)

		if matches == nil {
			errors = append(errors, "Meaning not parsed: "+word)
			continue
		}

		m := strings.Trim(matches[1], defaultWhitespace)
		p := strings.Trim(matches[3], defaultWhitespace)

		meanings = append(meanings, Meaning{m, p})
	}

	return meanings, errors
}

type DefaultWord struct {
	Id       string           `bson:"_id,omitempty" json:"_id,omitempty"`
	German   string           `bson:"german" json:"german,omitempty"`
	English  []Meaning        `bson:"english" json:"english,omitempty"`
	Third    []Meaning        `bson:"third" json:"third,omitempty"`
	Category string           `bson:"category" json:"category,omitempty"`
	User     string           `bson:"user" json:"user,omitempty"`
	Learned  time.Time        `bson:"learned" json:"learned,omitempty"`
	Score    int              `bson:"score" json:"score,omitempty"`
	Tags     []string         `bson:"tags" json:"tags,omitempty"`
	Errors   []string         `bson:"errors" json:"errors,omitempty"`
	Scores   []*general.Score `bson:"scores" json:"scores,omitempty"`
}

func NewDefaultWord(german, english, third, category, user, learned, score, tags string, errors []string) DefaultWord {
	englishMeanings, errors := NewMeanings(english, errors)
	thirdMeanings, errors := NewMeanings(third, errors)

	scoreParsed, err := strconv.ParseInt(score, 0, 0)
	if err != nil || scoreParsed < 1 || scoreParsed > 10 {
		scoreParsed = 5
	}

	learnedParsed := util.ParseTimeNow(learnedForm, learned)

	return DefaultWord{
		"",
		german,
		englishMeanings,
		thirdMeanings,
		category,
		user,
		learnedParsed,
		int(scoreParsed),
		util.TrimSplit(tags, tagSeparator),
		errors,
		[]*general.Score{},
	}
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

func (w *DefaultWord) GetScore() int {
	return w.Score
}

func (w *DefaultWord) GetUser() string {
	return w.User
}

func (w *DefaultWord) GetLearned() time.Time {
	return w.Learned
}

func (w *DefaultWord) GetErrors() []string {
	return w.Errors
}

func (w *DefaultWord) AddScore(score *general.Score) {
	w.Scores = append(w.Scores, score)
}

func (w *DefaultWord) GetScores() []*general.Score {
	return w.Scores
}

func (w *DefaultWord) GetId() string {
	return w.Id
}

type Any struct {
	DefaultWord `bson:"word" json:"word,omitempty"`
}

func NewAny(german, english, third, category, user, learned, score, tags string, errors []string) *Any {
	d := NewDefaultWord(german, english, third, category, user, learned, score, tags, errors)

	return &Any{d}
}
