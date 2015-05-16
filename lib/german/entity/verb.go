package entity

import (
	"errors"
	"regexp"
	"strings"

	germanUtil "github.com/peteraba/d5/lib/german/util"
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

type PersonalPronoun string

const (
	S1 PersonalPronoun = "S1"
	S2                 = "S2"
	S3                 = "S3"
	P1                 = "P1"
	P2                 = "P2"
	P3                 = "P3"
)

type Tense string

const (
	Present        Tense = "Present"
	Preterite            = "Preterite"
	PastParticiple       = "Past Participle"
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
		german, preterite, pastParticiple = main[0], main[1], main[2]
		break
	case 5:
		german, preterite, pastParticiple, du, er = main[0], main[1], main[2], main[3], main[4]
		break
	case 9:
		german, ich, du, er, wir, ihr, sie, preterite, pastParticiple = main[0], main[1], main[2], main[3], main[4], main[5], main[6], main[7], main[8]
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

func (v *Verb) getPresentStem() []string {
	p1s := []string{}

	for _, p1 := range v.P1 {
		if strings.HasSuffix(p1, "en") {
			p1s = append(p1s, strings.TrimSuffix(p1, "en"))
		} else if strings.HasSuffix(p1, "n") {
			p1s = append(p1s, strings.TrimSuffix(p1, "n"))
		}
	}

	return p1s
}

func (v *Verb) GetPresentS1() []string {
	if len(v.S1) > 0 {
		return v.S1
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "e")
}

func (v *Verb) GetPresentS2() []string {
	if len(v.S2) > 0 {
		return v.S2
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "st")
}

func (v *Verb) GetPresentS3() []string {
	if len(v.S3) > 0 {
		return v.S3
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "t")
}

func (v *Verb) GetPresentP1() []string {
	return v.P1
}

func (v *Verb) GetPresentP2() []string {
	if len(v.P2) > 0 {
		return v.P2
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "t")
}

func (v *Verb) GetPresentP3() []string {
	if len(v.P3) > 0 {
		return v.P3
	}

	return v.P1
}

func (v *Verb) GetPreteriteS1() []string {
	if len(v.Preterite) > 0 {
		return v.Preterite
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "te")
}

func (v *Verb) GetPreteriteS2() []string {
	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "st")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "test")
}

func (v *Verb) GetPreteriteS3() []string {
	if len(v.Preterite) > 0 {
		return v.Preterite
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "te")
}

func (v *Verb) GetPreteriteP1() []string {
	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "en")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "ten")
}

func (v *Verb) GetPreteriteP2() []string {
	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "t")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "tet")
}

func (v *Verb) GetPreteriteP3() []string {
	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "en")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "ten")
}

func (v *Verb) GetVerbPreterite(pp PersonalPronoun) []string {
	switch pp {
	case S1:
		return v.GetPreteriteS1()
	case S2:
		return v.GetPreteriteS2()
	case S3:
		return v.GetPreteriteS3()
	case P1:
		return v.GetPreteriteP1()
	case P2:
		return v.GetPreteriteP2()
	}

	return v.GetPreteriteP3()
}

func (v *Verb) GetVerbPresent(pp PersonalPronoun) []string {
	switch pp {
	case S1:
		return v.GetPresentS1()
	case S2:
		return v.GetPresentS2()
	case S3:
		return v.GetPresentS3()
	case P1:
		return v.GetPresentP1()
	case P2:
		return v.GetPresentP2()
	}

	return v.GetPresentP3()
}

func (v *Verb) GetVerb(pp PersonalPronoun, tense Tense) []string {
	switch tense {
	case Preterite:
		return v.GetVerbPreterite(pp)
	case Present:
		return v.GetVerbPresent(pp)
	}

	return v.PastParticiple
}

func (v *Verb) GetSeparated(pp PersonalPronoun, tense Tense) [][2]string {
	var (
		result       = [][2]string{}
		resultItem   [2]string
		nonSeparated []string
	)

	nonSeparated = v.GetVerb(pp, tense)

	for _, word := range nonSeparated {
		resultItem = [2]string{word, ""}
		if v.Prefix.Separable && v.Prefix.Prefix != "" && tense != PastParticiple {
			resultItem[0] = strings.TrimLeft(word, v.Prefix.Prefix)
			resultItem[1] = v.Prefix.Prefix
		}

		result = append(result, resultItem)
	}

	return result
}
