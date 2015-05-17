package entity

import (
	"regexp"

	"github.com/peteraba/d5/lib/german/dict"
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
