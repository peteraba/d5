package entity

import (
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
	Preposition string `bson:"prep" json:"prep,omitempty"`
	Case        Case   `bson:"case" json:"case,omitempty"`
}

func NewArgument(word string) Argument {
	matches := ArgumentRegexp.FindStringSubmatch(word)

	p := strings.Trim(matches[1], defaultWhitespace)
	c := strings.Trim(matches[3], defaultWhitespace)

	switch c {
	case "A":
		return Argument{p, CaseAcusative}
	case "N":
		return Argument{p, CaseNominative}
	case "D":
		return Argument{p, CaseDative}
	}

	return Argument{p, CaseGenitive}
}

func NewArguments(allArguments string) []Argument {
	arguments := []Argument{}

	allArguments = strings.TrimLeft(allArguments, argumentSeparator)

	for _, word := range util.TrimSplit(allArguments, argumentSeparator) {
		arguments = append(arguments, NewArgument(word))
	}

	return arguments
}

func parseArguments(rawArguments string) (Reflexive, []Argument, []string) {
	var (
		reflexive = ReflexiveWithout
		arguments = NewArguments(rawArguments)
		errors    = []string{}
	)

	if len(arguments) == 0 {
		return ReflexiveWithout, arguments, errors
	}

	if arguments[0].Preposition == "sich" {
		sich := arguments[0]
		arguments = arguments[1:]

		switch sich.Case {
		case "A":
			reflexive = ReflexiveAcusative
		case "D":
			reflexive = ReflexiveDative
		default:
			errors = append(errors, "Reflexive definition is invalid")
		}

		return reflexive, arguments, errors

	}

	return ReflexiveWithout, arguments, errors
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
	Prefix    string `bson:"prefix" json:"prefix,omitempty"`
	Separable bool   `bson:"separable" json:"separable,omitempty"`
}

// array of maps of word to exceptions
var separablePrefixes = [][]string{
	// top prio: causes hervor to be checked before her
	[]string{
		"auseinander",
		"entgegen",
		"entlang",
		"entzwei",
		"gegenüber",
		"gles1",
		"herbei",
		"herein",
		"herüber",
		"herunter",
		"hervor",
		"herauf",
		"heraus",
		"hinauf",
		"hinaus",
		"hinein",
		"hinterher",
		"hinunter",
		"hinweg",
		"nebenher",
		"nieder",
		"voraus",
		"vorbei",
		"vorüber",
		"vorweg",
		"zurecht",
		"zurück",
		"zusammen",
		"zwischen",
	},
	// moderate prio: causes herab to check before her
	[]string{
		"dabei",
		"daran",
		"s2rch",
		"empor",
		"fehl",
		"fest",
		"fort",
		"frei",
		"heim",
		"herab",
		"heran",
		"herum",
		"hinab",
		"hinzu",
		"hoch",
		"nach",
		"statt",
		"voran",
	},
	// low prio: causes vor to checked after vorher
	[]string{
		"an",
		"auf",
		"aus",
		"bei",
		"da",
		"dar",
		"ein",
		"her",
		"hin",
		"los",
		"mit",
		"vor",
		"weg",
		"zu",
	},
}

var separablePrefixExceptions = map[string][]string{
	"fehl": []string{"fehlen"},
}

// array of maps of word to exceptions
var unseparablePrefixes = [][]string{
	[]string{
		"be",
		"bei",
		"emp",
		"ent",
		"er",
		"ge",
		"miss",
		"ver",
		"voll",
		"zer",
	},
}

func NewPrefix(german string) Prefix {
	idx := strings.Index(german, "|")

	if idx > -1 {
		return Prefix{german[:idx], true}
	}

	for _, prefixSet := range separablePrefixes {
	separablePrefixLoop:
		for _, prefix := range prefixSet {
			if exceptions, ok := separablePrefixExceptions[prefix]; ok {
				if strings.Contains(strings.Join(exceptions, ","), prefix) {
					continue separablePrefixLoop
				}
			}

			if strings.Index(german, prefix) == 0 {
				return Prefix{prefix, true}
			}
		}
	}

	for _, prefixSet := range unseparablePrefixes {
		for _, prefix := range prefixSet {
			if strings.Index(german, prefix) == 0 {
				return Prefix{prefix, false}
			}
		}
	}

	return Prefix{"", false}
}

type Verb struct {
	DefaultWord    `bson:"word" json:"word,omitempty"`
	Auxiliary      []Auxiliary `bson:"auxiliary" json:"auxiliary,omitempty"`
	Prefix         Prefix      `bson:"prefix" json:"prefix,omitempty"`
	Noun           string      `bson:"noun" json:"noun,omitempty"`
	Adjective      string      `bson:"adjective" json:"adjective,omitempty"`
	PastParticiple []string    `bson:"pastParticiple" json:"pastParticiple,omitempty"`
	Preterite      []string    `bson:"preterite" json:"preterite,omitempty"`
	S1             []string    `bson:"s1" json:"s1,omitempty"`
	S2             []string    `bson:"s2" json:"s2,omitempty"`
	S3             []string    `bson:"s3" json:"s3,omitempty"`
	P1             []string    `bson:"p1" json:"p1,omitempty"`
	P2             []string    `bson:"p2" json:"p2,omitempty"`
	P3             []string    `bson:"p3" json:"p3,omitempty"`
	Reflexive      Reflexive   `bson:"reflexive" json:"reflexive,omitempty"`
	Arguments      []Argument  `bson:"arguments" json:"arguments,omitempty"`
	Id             string      `bson:"_id,omitempty" json:"_id,omitempty"`
}

func extractNounAdjective(german string) (string, string, string) {
	words := strings.Split(german, wordSeparator)

	if len(words) < 2 {
		return german, "", ""
	}

	german = words[len(words)-1]

	nouns, adjectives := []string{}, []string{}

	for _, word := range words[0 : len(words)-1] {
		if strings.ToLower(word) == word {
			adjectives = append(adjectives, word)
		} else {
			nouns = append(nouns, word)
		}
	}

	return german, strings.Join(nouns, wordSeparator), strings.Join(adjectives, wordSeparator)
}

func NewVerbP1(german string) string {
	german = strings.Replace(german, "|", "", -1)

	words := util.TrimSplit(german, wordSeparator)

	return words[len(words)-1]
}

func NewVerb(auxiliary, german, english, third, user, learned, score, tags string) *Verb {
	pastParticiple, preterite, s1, s2, s3, p1, p2, p3 := "", "", "", "", "", "", "", ""

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
		german, preterite, pastParticiple, s2, s3 = main[0], main[1], main[2], main[3], main[4]
		break
	case 9:
		german, s1, s2, s3, p1, p2, p3, preterite, pastParticiple = main[0], main[1], main[2], main[3], main[4], main[5], main[6], main[7], main[8]
		break
	default:
		return nil
	}

	sich, arguments, errors := parseArguments(matches[2])

	german, noun, adjective := extractNounAdjective(german)

	prefix := NewPrefix(german)

	german = strings.Replace(german, "|", "", -1)

	if p1 == "" {
		p1 = NewVerbP1(german)
	}

	return &Verb{
		NewDefaultWord(german, english, third, "verb", user, learned, score, tags, errors),
		NewAuxiliary(util.TrimSplit(auxiliary, alternativeSeparator)),
		prefix,
		noun,
		adjective,
		util.TrimSplit(pastParticiple, alternativeSeparator),
		util.TrimSplit(preterite, alternativeSeparator),
		util.TrimSplit(s1, alternativeSeparator),
		util.TrimSplit(s2, alternativeSeparator),
		util.TrimSplit(s3, alternativeSeparator),
		util.TrimSplit(p1, alternativeSeparator),
		util.TrimSplit(p2, alternativeSeparator),
		util.TrimSplit(p3, alternativeSeparator),
		sich,
		arguments,
		"",
	}
}

func (v *Verb) GetId() string {
	return v.Id
}

func (v *Verb) SetId(id string) {
	v.Id = id
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
	if len(v.S1) == 1 && v.S1[0] == "-" {
		return []string{"-"}
	}

	if len(v.Preterite) > 0 {
		return v.Preterite
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "te")
}

func (v *Verb) GetPreteriteS2() []string {
	if len(v.S2) == 1 && v.S2[0] == "-" {
		return []string{"-"}
	}

	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "st")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "test")
}

func (v *Verb) GetPreteriteS3() []string {
	if len(v.S3) == 1 && v.S3[0] == "-" {
		return []string{"-"}
	}

	if len(v.Preterite) > 0 {
		return v.Preterite
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "te")
}

func (v *Verb) GetPreteriteP1() []string {
	if len(v.P1) == 1 && v.P1[0] == "-" {
		return []string{"-"}
	}

	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "en")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "ten")
}

func (v *Verb) GetPreteriteP2() []string {
	if len(v.P2) == 1 && v.P2[0] == "-" {
		return []string{"-"}
	}

	if len(v.Preterite) > 0 {
		return germanUtil.SliceAppend(v.Preterite, "t")
	}

	return germanUtil.SliceAppend(v.getPresentStem(), "tet")
}

func (v *Verb) GetPreteriteP3() []string {
	if len(v.P3) == 1 && v.P3[0] == "-" {
		return []string{"-"}
	}

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
