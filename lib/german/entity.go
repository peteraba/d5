package german

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/peteraba/d5/lib/util"
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

type Auxiliary string

const (
	Sein  Auxiliary = "s"
	Haben           = "h"
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
	//  ([^(]*)                           -- match any string that is not a open parantheses character
	//          ?                         -- match optional space character
	//           (                        -- start of parantheses matching
	//            [(]                       -- match open parantheses character
	//               ([NADG])               -- match a case notion character  <-- N: nominative, A: acusative, D: dative, G: genitive
	//                       [)]            -- match close parantheses character
	//                          )?          -- end of parantheses matching
	//                             *      -- match optional spaces
	//                              $     -- match end of string
	ArgumentRegexp = regexp.MustCompile("^([^(]*) ?([(]([NADG])[)])? *$")

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

	// Noun:
	// ^                                                                                           -- match beginning of string
	//  ([A-ZÄÖÜ][A-ZÄÖÜßa-zäöü ]+)                                                                -- match noun in singular, must start with a capital
	//                             ,                                                               -- match a comma
	//                              ([A-ZÄÖÜa-zäöü~⍨ -]*)                                          -- match plural part, can be an extension only starting with a ⍨, ~
	//                                                     (,([A-ZÄÖÜßa-zäöü~⍨ ]*()?               -- match optional genitive, can be an extension
	//                                                                              ([(]pl[)])     -- match plural only note
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
	GetErrors() []string
}

type Argument struct {
	Preposition string `bson:"prep" json:"prep"`
	Case        string `bson:"case" json:"case"`
}

func NewArguments(allArguments string) []Argument {
	arguments := []Argument{}

	allArguments = strings.TrimLeft(allArguments, argumentSeparator)

	for _, word := range util.TrimSplit(allArguments, argumentSeparator) {
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

func parseArguments(rawArguments string) (Reflexive, []Argument, error) {
	arguments := NewArguments(rawArguments)

	if len(arguments) == 0 {
		return ReflexiveWithout, arguments, nil
	}

	if arguments[0].Preposition == "sich" {
		sich := arguments[0]
		arguments = arguments[1:]

		switch sich.Case {
		case "A":
			return ReflexiveAcusative, arguments, nil
			break
		case "B":
			return ReflexiveDative, arguments, nil
			break
		}

		return ReflexiveWithout, arguments, errors.New("Reflexive definition is invalid")

	}

	return ReflexiveWithout, arguments, nil
}

type Meaning struct {
	Main        string `bson:"main" json:"main"`
	Parantheses string `bson:"parantheses" json:"parantheses"`
}

func NewMeanings(allMeanings string) []Meaning {
	meanings := []Meaning{}

	for _, word := range util.TrimSplit(allMeanings, meaningSeparator) {
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

func NewAuxiliary(auxiliaries []string) []Auxiliary {
	var result []Auxiliary

	for _, auxiliary := range auxiliaries {
		switch auxiliary {
		case "h":
			result = append(result, Haben)
			break
		case "s":
			result = append(result, Sein)
			break
		}
	}

	return result
}

type Prefix struct {
	Prefix    string `bson:"prefix" json:"prefix"`
	Separable bool   `bson:"separable" json:"separable"`
}

// array of maps of word to exceptions
var separablePrefixes = []map[string][]string{
	// top prio: causes hervor to be checked before her
	map[string][]string{
		"auseinander": []string{},
		"entgegen":    []string{},
		"entlang":     []string{},
		"entzwei":     []string{},
		"gegenüber":   []string{},
		"gleich":      []string{},
		"herbei":      []string{},
		"herein":      []string{},
		"herüber":     []string{},
		"herunter":    []string{},
		"hervor":      []string{},
		"herauf":      []string{},
		"heraus":      []string{},
		"hinauf":      []string{},
		"hinaus":      []string{},
		"hinein":      []string{},
		"hinterher":   []string{},
		"hinunter":    []string{},
		"hinweg":      []string{},
		"nebenher":    []string{},
		"nieder":      []string{},
		"voraus":      []string{},
		"vorbei":      []string{},
		"vorüber":     []string{},
		"vorweg":      []string{},
		"zurecht":     []string{},
		"zurück":      []string{},
		"zusammen":    []string{},
		"zwischen":    []string{},
	},
	// moderate prio: causes herab to check before her
	map[string][]string{
		"dabei": []string{},
		"daran": []string{},
		"durch": []string{},
		"empor": []string{},
		"fehl":  []string{"fehlen"},
		"fest":  []string{},
		"fort":  []string{},
		"frei":  []string{},
		"heim":  []string{},
		"herab": []string{},
		"heran": []string{},
		"herum": []string{},
		"hinab": []string{},
		"hinzu": []string{},
		"hoch":  []string{},
		"nach":  []string{},
		"statt": []string{},
		"voran": []string{},
	},
	// low prio: causes vor to checked after vorher
	map[string][]string{
		"an":  []string{},
		"auf": []string{},
		"aus": []string{},
		"bei": []string{},
		"da":  []string{},
		"dar": []string{},
		"ein": []string{},
		"her": []string{},
		"hin": []string{},
		"los": []string{},
		"mit": []string{},
		"vor": []string{},
		"weg": []string{},
		"zu":  []string{},
	},
}

// array of maps of word to exceptions
var unseparablePrefixes = []map[string][]string{
	map[string][]string{
		"be":   []string{},
		"bei":  []string{},
		"emp":  []string{},
		"ent":  []string{},
		"er":   []string{},
		"ge":   []string{},
		"miss": []string{},
		"ver":  []string{},
		"voll": []string{},
		"zer":  []string{},
	},
}

func NewPrefix(german string) Prefix {
	idx := strings.Index(german, "|")

	if idx > -1 {
		return Prefix{german[:idx], true}
	}

	for _, prefixSet := range separablePrefixes {
	separablePrefixLoop:
		for prefix, exceptions := range prefixSet {
			for _, exception := range exceptions {
				if exception == german {
					continue separablePrefixLoop
				}
			}

			if strings.Index(german, prefix) == 0 {
				return Prefix{prefix, true}
			}
		}
	}

	for _, prefixSet := range unseparablePrefixes {
	unseparablePrefixLoop:
		for prefix, exceptions := range prefixSet {
			for _, exception := range exceptions {
				if exception == german {
					continue unseparablePrefixLoop
				}
			}

			if strings.Index(german, prefix) == 0 {
				return Prefix{prefix, false}
			}
		}
	}

	return Prefix{"", false}
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
	Errors   []string  `bson:"errors", json:"errors"`
}

func NewDefaultWord(german, english, third, category, user, learned, score, tags string, errors []string) DefaultWord {
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
		util.TrimSplit(tags, tagSeparator),
		errors,
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

type Any struct {
	DefaultWord `bson:"word" json:"word"`
}

func NewAny(german, english, third, category, user, learned, score, tags string, errors []string) *Any {
	d := NewDefaultWord(german, english, third, category, user, learned, score, tags, errors)

	return &Any{d}
}

type Adjective struct {
	DefaultWord `bson:"word" json:"word"`
	Comparative []string `bson:"comparative" json:"comparative"`
	Superlative []string `bson:"superlative" json:"superlative"`
}

func NewAdjective(german, english, third, user, learned, score, tags string) *Adjective {
	adjectiveParts := util.TrimSplit(german, conjugationSeparator)

	if len(adjectiveParts) < 1 {
		return nil
	}

	errors := []string{}
	comparative := []string{}
	superlative := []string{}

	german = adjectiveParts[0]

	if len(adjectiveParts) > 1 {
		comparative = util.TrimSplit(adjectiveParts[1], alternativeSeparator)
	}
	if len(adjectiveParts) > 2 {
		superlative = util.TrimSplit(adjectiveParts[2], alternativeSeparator)
	}

	return &Adjective{
		NewDefaultWord(german, english, third, "adjective", user, learned, score, tags, errors),
		comparative,
		superlative,
	}
}

type Noun struct {
	DefaultWord  `bson:"word" json:"word"`
	Articles     []Article `bson:"article" json:"article"`
	Plural       []string  `bson:"plural" json:"plural"`
	Genitive     []string  `bson:"genitive" json:"genitive"`
	IsPluralOnly bool      `bson:"plural_only" json:"plural_only"`
}

func NewNoun(articles, german, english, third, user, learned, score, tags string) *Noun {
	matches := NounRegexp.FindStringSubmatch(german)

	if len(matches) < 5 {
		return nil
	}

	errors := []string{}

	articleList := []Article{}
	for _, article := range util.TrimSplit(articles, alternativeSeparator) {
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

	german = matches[1]

	return &Noun{
		NewDefaultWord(german, english, third, "noun", user, learned, score, tags, errors),
		articleList,
		util.TrimSplit(matches[2], alternativeSeparator),
		util.TrimSplit(matches[4], alternativeSeparator),
		matches[5] == "(pl)",
	}
}

type Verb struct {
	DefaultWord    `bson:"word" json:"word"`
	Auxiliary      []Auxiliary `bson:"auxiliary" json:"auxiliary"`
	Prefix         Prefix      `bson:"prefix" json:"prefix"`
	Noun           string      `bson:"noun" json:"noun"`
	Adjective      string      `bson:"adjective" json:"adjective"`
	PastParticiple []string    `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      []string    `bson:"preterite" json:"preterite"`
	S1             []string    `bson:"s1" json:"s1"`
	S2             []string    `bson:"s2" json:"s2"`
	S3             []string    `bson:"s3" json:"s3"`
	P1             []string    `bson:"p1" json:"p1"`
	P2             []string    `bson:"p2" json:"p2"`
	P3             []string    `bson:"p3" json:"p3"`
	Reflexive      Reflexive   `bson:"reflexive" json:"reflexive"`
	Arguments      []Argument  `bson:"arguments" json:"arguments"`
}

func NewVerbNoun(german string) string {
	words := util.TrimSplit(german, conjugationSeparator)

	if len(words) == 0 {
		return ""
	}

	nouns := []string{}
	for _, word := range words[1:] {
		if strings.ToLower(word) != word {
			nouns = append(nouns, word)
		}
	}

	return strings.Join(nouns, " ")
}

func NewVerbAdjective(german string) string {
	words := util.TrimSplit(german, conjugationSeparator)

	if len(words) == 0 {
		return ""
	}

	adjectives := []string{}
	for _, word := range words[1:] {
		if strings.ToLower(word) == word {
			adjectives = append(adjectives, word)
		}
	}

	return strings.Join(adjectives, " ")
}

func getGermanNounAdjective(german string) (string, string, string) {
	german = strings.Replace(german, "|", "", -1)

	noun := NewVerbNoun(german)
	adjective := NewVerbAdjective(german)

	if noun != "" {
		german = noun + " " + german
	}

	if adjective != "" {
		german = adjective + " " + german
	}

	return german, noun, adjective
}

func NewVerb(auxiliary, german, english, third, user, learned, score, tags string) *Verb {
	pastParticiple, preterite, ich, du, er, wir, ihr, sie := "", "", "", "", "", "", "", ""

	matches := VerbRegexp.FindStringSubmatch(german)
	if len(matches) < 3 {
		return nil
	}

	errors := []string{}

	main := util.TrimSplit(matches[1], conjugationSeparator)
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

	prefix := NewPrefix(german)

	german, noun, adjective := getGermanNounAdjective(german)

	return &Verb{
		NewDefaultWord(german, english, third, "verb", user, learned, score, tags, errors),
		NewAuxiliary(util.TrimSplit(auxiliary, alternativeSeparator)),
		prefix,
		noun,
		adjective,
		util.TrimSplit(pastParticiple, alternativeSeparator),
		util.TrimSplit(preterite, alternativeSeparator),
		util.TrimSplit(ich, alternativeSeparator),
		util.TrimSplit(du, alternativeSeparator),
		util.TrimSplit(er, alternativeSeparator),
		util.TrimSplit(wir, alternativeSeparator),
		util.TrimSplit(ihr, alternativeSeparator),
		util.TrimSplit(sie, alternativeSeparator),
		sich,
		arguments,
	}
}
