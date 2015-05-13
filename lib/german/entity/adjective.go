package entity

import (
	"regexp"

	"github.com/peteraba/d5/lib/german/dict"
	"github.com/peteraba/d5/lib/util"
)

const (
	comparativeJoin = ","
	superlativeJoin = ","
)

var (
	// Adjective:
	// ^                                                       -- match beginning of string
	//  ([a-zäöüß]+)                                           -- match adjective
	//              (,([a-zäöüß~⍨-]*))?                        -- match optional comparative, can be an extension only starting with a ⍨, ~
	//                                 (,([a-zäöüß~⍨-]*))?     -- match optional superlative, can be an extension
	//                                                    $    -- match end of string
	AdjectiveRegexp = regexp.MustCompile("^([a-zäöüß]+)(,([a-zäöüß~⍨-]*))?(,([a-zäöüß~⍨-]*))?$")
)

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

func (a *Adjective) GetComparative() []string {
	result := []string{}
	for _, comparative := range a.Comparative {
		result = append(result, dict.Decline(a.German, comparative))
	}

	return result
}

func (a *Adjective) GetComparativeString(maxCount int) string {
	raw := a.GetComparative()

	return util.JoinLimited(raw, comparativeJoin, maxCount)
}

func (a *Adjective) GetSuperlative() []string {
	result := []string{}
	for _, superlative := range a.Superlative {
		result = append(result, dict.Decline(a.German, superlative))
	}

	return result
}

func (a *Adjective) GetSuperlativeString(maxCount int) string {
	raw := a.GetSuperlative()

	return util.JoinLimited(raw, genitiveJoin, maxCount)
}
