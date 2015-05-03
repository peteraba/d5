package german

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Reflexive string

const (
	ReflexiveWithout   Reflexive = ""
	ReflexiveAcusative           = "A"
	ReflexiveDative              = "D"
)

type Case string

const (
	CaseNominative Case = "N"
	CaseAcusative       = "A"
	CaseDative          = "D"
	CaseGenitive        = "G"
)

type Article string

const (
	Der Article = "r"
	Die         = "e"
	Das         = "s"
)

const learnedForm = "2006-01-02"

const (
	alternativeSeparator = "/"
	conjugationSeparator = ","
	argumentSeparator    = "+"
	meaningSeparator     = ";"
	synonimSeparator     = ","
	tagSeparator         = ","
	defaultWhitespace    = "\t\n\f\r "
)

var (
	// Article:
	// ^                      -- match beginning of string
	//  ([res])               -- match first article notion <-- r: der, e: die, s: das
	//         (/([res]))?    -- match optional second article notion, following a / sign
	//                    $   -- match end of string
	ArticleRegexp = regexp.MustCompile("^([res])(/([res]))?$")

	// Auxiliary:
	// ^                      -- match beginning of string
	//  ([sh])                -- match first auxiliary notion <-- s: sein, h: haben
	//        (/([hs]))?      -- match optional second auxiliary notion, followin a / sign
	//                  $     -- match end of string
	AuxiliaryRegexp = regexp.MustCompile("^([sh])(/([hs]))?$")

	// Argument:
	// ^                                  -- match beginning of string
	//  ([^(]*)                           -- match any string that is not a open paranthases character
	//          ?                         -- match optional space character
	//           (                        -- start of paranthases matching
	//            [(]                       -- match open paranthases character
	//               ([NADG])               -- match a case notion character  <-- N: nominative, A: acusative, D: dative, G: genitive
	//                       [)]            -- match close paranthases character
	//                          )?          -- end of paranthases matching
	//                             *      -- match optional spaces
	//                              $     -- match end of string
	ArgumentRegexp = regexp.MustCompile("^([^(]*) ?([(]([NADG])[)])? *$")

	// Meaning:
	// ^                                -- match beginning of string
	//  ([^(]*)                         -- match any string that is not a open paranthases character
	//          ?                       -- match optional space character
	//           (                      -- start of paranthases matching
	//            [(]                     -- match open paranthases character
	//               ([^)]*)              -- match paranthases content
	//                      [)]           -- match close paranthases character
	//                         )?         -- end of paranthases matching
	//                           *      -- match optional spaces
	//                            $     -- match end of string
	MeaningRegexp = regexp.MustCompile("^([^(]*) ?([(]([^)]*)[)])? *$")

	// Noun:
	// ^                                                                                          -- match beginning of string
	//  ([A-ZÄÖÜ][A-ZÄÖÜßa-zäöü ]+)                                                               -- match noun in singular, must start with a capital
	//                             ,                                                              -- match a comma
	//                              ([A-ZÄÖÜa-zäöü~⍨ -]*)                                         -- match plural part, can be an extension only starting with a ⍨, ~
	//                                                     (,([A-ZÄÖÜßa-zäöü~⍨ ]*()?               -- match optional genitive, can be an extension
	//                                                                              ([(]pl[)])
	//                                                                                        $    -- match end of string
	NounRegexp = regexp.MustCompile("^([A-ZÄÖÜ][A-ZÄÖÜßa-zäöü ]+),([A-ZÄÖÜa-zäöü~⍨/ -]*)(,([A-ZÄÖÜßa-zäöü~⍨/ ]*))?([(]pl[)])?$")

	// Adjective:
	// ^                                                       -- match beginning of string
	//  ([a-zäöüß]+)                                           -- match adjective
	//              (,([a-zäöüß~⍨-]*))?                        -- match optional comparative, can be an extension only starting with a ⍨, ~
	//                                 (,([a-zäöüß~⍨-]*))?     -- match optional superlative, can be an extension
	//                                                    $    -- match end of string
	AdjectiveRegexp = regexp.MustCompile("^([a-zäöüß]+)(,([a-zäöüß~⍨-]*))?(,([a-zäöüß~⍨-]*))?$")

	// Verb:
	// ^                                                 -- match beginning of string
	//  ([A-ZÄÖÜßa-zäöü, ]+)                             -- match verb
	//                     ([A-ZÄÖÜßa-zäöü+() ]*)?       -- match extension(s), separated by plus signs
	//                                            $      -- match end of string
	VerbRegexp = regexp.MustCompile("^([A-ZÄÖÜßa-zäöü|, ]+)([A-ZÄÖÜßa-zäöü+() ]*)?$")

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
	IsOk() bool
}

type Argument struct {
	Preposition string `bson:"prep" json:"prep"`
	Case        string `bson:"case" json:"case"`
}

func NewArgument(allArguments string) []Argument {
	arguments := []Argument{}

	allArguments = strings.TrimLeft(allArguments, argumentSeparator)

	for _, word := range TrimSplit(allArguments, argumentSeparator) {
		matches := ArgumentRegexp.FindStringSubmatch(word)
		if len(matches) < 3 {
			continue
		}

		p := strings.Trim(matches[1], defaultWhitespace)
		c := strings.Trim(matches[3], defaultWhitespace)

		arguments = append(arguments, Argument{p, c})
	}

	return arguments
}

type Meaning struct {
	Main        string `bson:"main" json:"main"`
	Paranthases string `bson:"paranthases" json:"paranthases"`
}

func NewMeanings(allMeanings string) []Meaning {
	meanings := []Meaning{}

	for _, word := range TrimSplit(allMeanings, meaningSeparator) {
		matches := MeaningRegexp.FindStringSubmatch(word)
		if len(matches) < 3 {
			continue
		}

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
	Ok       bool      `bson:"ok", json:"ok"`
}

func NewDefaultWord(german, english, third, category, user, learned, score, tags string) DefaultWord {
	englishMeanings, thirdMeanings := NewMeanings(english), NewMeanings(third)

	scoreParsed, err := strconv.ParseInt(score, 0, 0)
	if err != nil || scoreParsed < 1 || scoreParsed > 10 {
		scoreParsed = 5
	}

	learnedParsed, err := time.Parse(learnedForm, learned)
	if err != nil {
		learnedParsed = time.Now()
	}

	return DefaultWord{
		german,
		englishMeanings,
		thirdMeanings,
		category,
		user,
		learnedParsed,
		int(scoreParsed),
		TrimSplit(tags, tagSeparator),
		true,
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

func (w *DefaultWord) IsOk() bool {
	return w.Ok
}

type Any struct {
	DefaultWord `bson:"word" json:"word"`
}

func NewAny(german, english, third, category, user, learned, score, tags string, ok bool) *Any {
	d := NewDefaultWord(german, english, third, category, user, learned, score, tags)

	d.Ok = ok

	return &Any{d}
}

type Adjective struct {
	DefaultWord `bson:"word" json:"word"`
	Comparative []string `bson:"comparative" json:"comparative"`
	Superlative []string `bson:"superlative" json:"superlative"`
}

func NewAdjective(german, english, third, user, learned, score, tags string) *Adjective {
	adjectiveParts := TrimSplit(german, conjugationSeparator)

	comparative := []string{}
	superlative := []string{}

	if len(adjectiveParts) > 1 {
		comparative = TrimSplit(adjectiveParts[1], alternativeSeparator)
	}
	if len(adjectiveParts) > 2 {
		superlative = TrimSplit(adjectiveParts[2], alternativeSeparator)
	}

	return &Adjective{
		NewDefaultWord(german, english, third, "adjective", user, learned, score, tags),
		comparative,
		superlative,
	}
}

type Noun struct {
	DefaultWord  `bson:"word" json:"word"`
	Articles     []Article `bson:"article" json:"article"`
	Plural       []string  `bson:"plural" json:"plural"`
	IsPluralOnly bool      `bson:"plural_only" json:"plural_only"`
}

func NewNoun(articles, german, english, third, user, learned, score, tags string) *Noun {
	matches := NounRegexp.FindStringSubmatch(german)

	if len(matches) < 5 {
		return nil
	}

	articleList := []Article{}
	for _, article := range TrimSplit(articles, alternativeSeparator) {
		switch article {
		case "r":
			articleList = append(articleList, Der)
			break
		case "e":
			articleList = append(articleList, Die)
			break
		case "s":
			articleList = append(articleList, Das)
			break
		default:
			return nil
		}
	}

	return &Noun{
		NewDefaultWord(german, english, third, "noun", user, learned, score, tags),
		articleList,
		TrimSplit(matches[3], alternativeSeparator),
		matches[4] == "(pl)",
	}
}

type Verb struct {
	DefaultWord    `bson:"word" json:"word"`
	Auxiliary      []string  `bson:"auxiliary" json:"auxiliary"`
	Noun           string    `bson:"noun" json:"noun"`
	Adjective      string    `bson:"adjective" json:"adjective"`
	PastParticiple []string  `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      []string  `bson:"preterite" json:"preterite"`
	S1             []string  `bson:"s1" json:"s1"`
	S2             []string  `bson:"s2" json:"s2"`
	S3             []string  `bson:"s3" json:"s3"`
	P1             []string  `bson:"p1" json:"p1"`
	P2             []string  `bson:"p2" json:"p2"`
	P3             []string  `bson:"p3" json:"p3"`
	Reflexive      Reflexive `bson:"reflexive" json:"reflexive"`
	Arguments      []string  `bson:"arguments" json:"arguments"`
}

func NewVerb(auxiliary, german, english, third, user, learned, score, tags string) *Verb {
	pastParticiple, preterite, ich, du, er, wir, ihr, sie := "", "", "", "", "", "", "", ""

	matches := VerbRegexp.FindStringSubmatch(german)
	if len(matches) < 3 {
		return nil
	}

	main := TrimSplit(matches[1], conjugationSeparator)
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

	sich, arguments, err := parseArguments(matches[2])
	if err != nil {
		return nil
	}

	return &Verb{
		NewDefaultWord(german, english, third, "verb", user, learned, score, tags),
		TrimSplit(auxiliary, alternativeSeparator),
		"",
		"",
		TrimSplit(pastParticiple, alternativeSeparator),
		TrimSplit(preterite, alternativeSeparator),
		TrimSplit(ich, alternativeSeparator),
		TrimSplit(du, alternativeSeparator),
		TrimSplit(er, alternativeSeparator),
		TrimSplit(wir, alternativeSeparator),
		TrimSplit(ihr, alternativeSeparator),
		TrimSplit(sie, alternativeSeparator),
		sich,
		arguments,
	}
}

func parseArguments(rawArguments string) (Reflexive, []string, error) {
	if rawArguments == "" {
		return ReflexiveWithout, []string{}, nil
	}
	arguments := TrimSplit(rawArguments, argumentSeparator)

	if strings.Contains(arguments[0], "sich (A)") {
		return ReflexiveAcusative, arguments[1:], nil
	}

	if strings.Contains(arguments[0], "sich (D)") {
		return ReflexiveDative, arguments[1:], nil
	}

	if strings.Contains(arguments[0], "sich") {
		return ReflexiveWithout, []string{}, errors.New("Reflexive definition is invalid")
	}

	return ReflexiveWithout, []string{}, nil
}
