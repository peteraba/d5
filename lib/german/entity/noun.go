package entity

import (
	"regexp"
	"strings"

	"github.com/peteraba/d5/lib/german/dict"
	germanUtil "github.com/peteraba/d5/lib/german/util"
	"github.com/peteraba/d5/lib/util"
)

const (
	pluralJoin   = ", "
	genitiveJoin = ", "
)

var (
	// Noun:
	// ^                                                                                           -- match beginning of string
	//  ([A-ZÄÖÜ][A-ZÄÖÜßa-zäöü ]+)                                                                -- match noun in singular, must start with a capital
	//                             ,                                                               -- match a comma
	//                              ([A-ZÄÖÜa-zäöü~⍨ -]*)                                          -- match plural part, can be an extension only starting with a ⍨, ~
	//                                                     (,([A-ZÄÖÜßa-zäöü~⍨ ]*()?               -- match optional genitive, can be an extension
	//                                                                              ([(]pl[)])     -- match plural only note
	//                                                                                        $    -- match end of string
	NounRegexp = regexp.MustCompile("^([A-ZÄÖÜ][A-ZÄÖÜßa-zäöü -]+),([A-ZÄÖÜa-zäöü~⍨/ -]*)(,([A-ZÄÖÜßa-zäöü~⍨/ -]*))?([(]pl[)])?$")
)

type Noun struct {
	DefaultWord  `bson:"word" json:"word,omitempty"`
	Articles     []Article `bson:"article" json:"article,omitempty"`
	Plural       []string  `bson:"plural" json:"plural,omitempty"`
	Genitive     []string  `bson:"genitive" json:"genitive,omitempty"`
	IsPluralOnly bool      `bson:"plural_only" json:"plural_only,omitempty"`
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
		default:
			articleList = append(articleList, Das)
			break
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

func (n *Noun) GetPlurals() []string {
	if n.IsPluralOnly {
		return []string{n.German}
	}

	result := []string{}
	for _, pl := range n.Plural {
		result = append(result, dict.Decline(n.German, pl))
	}

	return result
}

func (n *Noun) GetPluralsString(maxCount int) string {
	raw := n.GetPlurals()

	return util.JoinLimited(raw, pluralJoin, maxCount)
}

func (n *Noun) GetGenitives() []string {
	result := []string{}
	for _, genitive := range n.Genitive {
		result = append(result, dict.Decline(n.German, genitive))
	}

	return result
}

func (n *Noun) GetGenitivesString(maxCount int) string {
	raw := n.GetGenitives()

	return util.JoinLimited(raw, genitiveJoin, maxCount)
}

// http://en.wikipedia.org/wiki/German_nouns#Declension_for_case
func (n *Noun) Decline(
	isPlural bool,
	nounCase Case,
) []string {
	result := []string{}

	// For plural nouns
	if isPlural {
		var lastChar string

		result = n.GetPlurals()

		if nounCase == CaseDative {
			for key, word := range result {
				lastChar = word[len(word)-1:]
				if lastChar == "n" || lastChar == "s" {
					continue
				}

				result[key] = germanUtil.AddSuffix(word, "n")
			}
		}

		// Generate plural
		return result
	}

	// Use provided data when present
	if nounCase == CaseGenitive && len(n.GetGenitives()) > 0 {
		return n.GetGenitives()
	}

	// I: Feminine nouns have the same form in all four cases.
	if n.Articles[0] == Die {
		result = append(result, n.German)

		return result
	}

	// III: The n-nouns take -(e)n for genitive, dative and accusative: this is used for masculine nouns ending with -e and a few others, mostly animate nouns.
	if n.IsWeak() || n.IsMixed() {
		if nounCase == CaseAcusative || nounCase == CaseDative || nounCase == CaseGenitive {
			ending := ""

			if !(n.IsMixed() && nounCase == CaseAcusative) {
				ending = "en"
				if germanUtil.IsVowel(n.German[len(n.German)-1:]) {
					ending = "n"
				}

				if nounCase == CaseGenitive && n.IsMixed() {
					ending += "s"
				}
			}

			result = append(result, n.German+ending)

			return result
		}
	}

	// II: Personal names, All neuter and most masculine nouns have genitive case '-(e)s' endings: normally '-es' if one syllable long, '-s' if more. This is related to using 's to show possession in English, e.g. 'The boy's book'. Traditionally the nouns in this group also add -e in the dative case, but this is now often ignored.
	if n.Articles[0] == Der || n.Articles[0] == Das {
		if nounCase == CaseDative {
			// Add optional ~e
			result = append(result, n.German)
			result = append(result, n.German+"e")

			return result
		}

		if nounCase == CaseGenitive {
			// Add s or es depending on syllable count
			var sAdded = germanUtil.AddSuffix(n.German, "s")

			result = append(result, sAdded)

			if germanUtil.CountSyllables(n.German) == 1 {
				if strings.LastIndex(sAdded, "es") != len(sAdded)-2 {
					result = append(result, germanUtil.AddSuffix(n.German, "es"))
				}
			}

			return result
		}
	}

	// IV: A few masculine n-nouns take (e)n for accusative and dative, and -(e)ns for genitive.
	// Todo

	result = append(result, n.GetGerman())

	// Generate single
	return result
}

var mixedNouns = []string{"Herz", "Buchstabe", "Gedanke", "Friede", "Glaube", "Wille"}

func (n *Noun) IsMixed() bool {
	return util.StringIn(n.German, mixedNouns)
}

var weekEndingExceptions = []string{"Käse"}
var weakEndings = []string{"ant", "ent", "ist", "aut", "at", "ad"}
var weakNouns = []string{
	"Herr", "Bauer", "Bär", "Held", "Mensch",
	"Nachbarn", "Pilot", "Idiot", "Architekt", "Prinz",
	"Präsident", "Philosph",
}

// http://germanforenglishspeakers.com/nouns/weak-nouns-the-n-declension/
func (n *Noun) IsWeak() bool {
	if len(n.Articles) > 1 || n.Articles[0] != Der {
		return false
	}

	if strings.HasSuffix(n.German, "e") && !util.StringIn(n.German, weekEndingExceptions) {
		return true
	}

	if util.HasSuffixAny(n.German, weakEndings) {
		return true
	}

	if util.StringIn(n.German, weakNouns) {
		return true
	}

	return false
}
