package entity

import (
	"regexp"
	"strings"

	"github.com/peteraba/d5/lib/german/dict"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2/bson"
)

const (
	comparativeJoin = ","
	superlativeJoin = ","
)

type Degree int

const (
	Positive    Degree = 0
	Comparative        = 1
	Superlative        = 2
)

type Declension int

const (
	Strong Declension = 0
	Weak              = 1
	Mixed             = 2
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
	DefaultWord `bson:"word" json:"word,omitempty"`
	Comparative []string      `bson:"comparative" json:"comparative,omitempty"`
	Superlative []string      `bson:"superlative" json:"superlative,omitempty"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
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
		"",
	}
}

func (a *Adjective) GetId() bson.ObjectId {
	return a.Id
}

func (a *Adjective) SetId(id bson.ObjectId) {
	a.Id = id
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

func (a *Adjective) Decline(
	degree Degree,
	declension Declension,
	nounArticle Article,
	isPlural bool,
	nounCase Case,
) []string {
	var (
		words  []string
		ending string
	)

	switch degree {
	case Positive:
		words = []string{a.GetGerman()}
		break
	case Comparative:
		words = a.GetComparative()
		break
	default:
		words = a.GetSuperlative()
	}

	switch declension {
	case Strong:
		ending = strongInflection(nounArticle, isPlural, nounCase)
		break
	case Weak:
		ending = weakInflection(nounArticle, isPlural, nounCase)
		break
	default:
		ending = mixedInlection(nounArticle, isPlural, nounCase)
	}

	ending = strings.TrimLeft(ending, "~")

	return util.SliceAppend(words, ending)
}

func strongInflection(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~er"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}

		return "~e"
	case CaseAcusative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}
		return "~e"
	case CaseDative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~em"
			case Die:
				return "~er"
			case Das:
				return "~em"
			}
		}
		return "~en"
	}

	//case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "~en"
		case Die:
			return "~er"
		case Das:
			return "~en"
		}
	}
	return "~er"
}

func mixedInlection(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~er"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}
		return "~en"
	case CaseAcusative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}
		return "~en"
	case CaseDative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~en"
			case Das:
				return "~en"
			}
		}
		return "~en"
	}

	//case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "~en"
		case Die:
			return "~en"
		case Das:
			return "~en"
		}
	}
	return "~en"
}

func weakInflection(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~e"
			case Die:
				return "~e"
			case Das:
				return "~e"
			}
		}
		return "~en"
	case CaseAcusative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~e"
			case Das:
				return "~e"
			}
		}
		return "~en"
	case CaseDative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~en"
			case Das:
				return "~en"
			}
		}
		return "~en"
	}

	//case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "~en"
		case Die:
			return "~en"
		case Das:
			return "~en"
		}
	}

	return "~en"
}
