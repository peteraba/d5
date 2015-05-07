package entity

import (
	"errors"
	"regexp"
	"strings"

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

type Auxiliary string

const (
	Sein  Auxiliary = "s"
	Haben           = "h"
)

const (
	argumentSeparator = "+"
)

var (
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

	// Verb:
	// ^                                                 -- match beginning of string
	//  ([A-ZÄÖÜßa-zäöü, ]+)                             -- match verb
	//                     ([A-ZÄÖÜßa-zäöü+() ]*)?       -- match extension(s), separated by plus signs
	//                                            $      -- match end of string
	VerbRegexp = regexp.MustCompile("^([A-ZÄÖÜßa-zäöü|,/ -]+)([A-ZÄÖÜßa-zäöü+()/ -]*)?$")
)

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

func extractNounAdjective(german string) (string, string) {
	lastIndex := strings.LastIndex(german, wordSeparator)
	if lastIndex == -1 {
		return "", ""
	}

	noun, adjective := "", ""

	extra := german[0:lastIndex]

	if strings.ToLower(extra) == extra {
		adjective = extra
	} else {
		noun = extra
	}

	return noun, adjective
}

func NewVerbWir(german string) string {
	german = strings.Replace(german, "|", "", -1)

	words := util.TrimSplit(german, wordSeparator)

	return words[len(words)-1]
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

	if wir == "" {
		wir = NewVerbWir(german)
	}

	german = strings.Replace(german, "|", "", -1)

	noun, adjective := extractNounAdjective(german)

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
