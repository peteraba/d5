package entity

import (
	"regexp"
	"strings"
)

type Article string

const (
	Der Article = "r"
	Die         = "e"
	Das         = "s"
)

var (
	// Article:
	// ^                      -- match beginning of string
	//  ([res])               -- match first article notion <-- r: der, e: die, s: das
	//         (/([res]))?    -- match optional second article notion, following a / sign
	//                    $   -- match end of string
	ArticleRegexp = regexp.MustCompile("^([res])(/([res]))?$")
)

func indefiniteEnding(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "~"
			case Die:
				return "~e"
			case Das:
				return "~"
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
				return "~"
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

	// Case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "~es"
		case Die:
			return "~er"
		case Das:
			return "~es"
		}
	}

	return "~er"
}

func definiteEnding(nounArticle Article, isPlural bool, nounCase Case) string {
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

	// Case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "~es"
		case Die:
			return "~er"
		case Das:
			return "~es"
		}
	}

	return "~er"
}

func IndefiniteArticle(word string, nounArticle Article, isPlural bool, nounCase Case) string {
	if word == "" {
		word = "ein"
	}

	if word == "ein" && isPlural {
		return ""
	}

	ending := indefiniteEnding(nounArticle, isPlural, nounCase)

	return strings.TrimRight(word, "e") + strings.TrimLeft(ending, "~")
}

func DefiniteArticle(word string, nounArticle Article, isPlural bool, nounCase Case) string {
	if word == "der" || word == "die" || word == "das" {
		return declineArticle(nounArticle, isPlural, nounCase)
	}

	ending := definiteEnding(nounArticle, isPlural, nounCase)

	return strings.TrimRight(word, "e") + strings.TrimLeft(ending, "~")
}

func declineArticle(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "der"
			case Die:
				return "die"
			case Das:
				return "das"
			}
		}
		return "die"
	case CaseAcusative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "den"
			case Die:
				return "die"
			case Das:
				return "das"
			}
		}
		return "die"
	case CaseDative:
		if !isPlural {
			switch nounArticle {
			case Der:
				return "dem"
			case Die:
				return "der"
			case Das:
				return "dem"
			}
		}
		return "den"
	}

	// Case CaseGenitive:
	if !isPlural {
		switch nounArticle {
		case Der:
			return "des"
		case Die:
			return "der"
		case Das:
			return "des"
		}
	}

	return "der"
}
