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
		if isPlural {
			return "~e"
		} else {
			switch nounArticle {
			case Der:
				return "~"
			case Die:
				return "~e"
			case Das:
				return "~"
			}
		}
		break
	case CaseAcusative:
		if isPlural {
			return "~e"
		} else {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~e"
			case Das:
				return "~"
			}
		}
		break
	case CaseDative:
		if isPlural {
			return "~en"
		} else {
			switch nounArticle {
			case Der:
				return "~em"
			case Die:
				return "~er"
			case Das:
				return "~em"
			}
		}
		break
	case CaseGenitive:
		if isPlural {
			return "~er"
		} else {
			switch nounArticle {
			case Der:
				return "~es"
			case Die:
				return "~er"
			case Das:
				return "~es"
			}
		}
		break
	}

	return ""
}

func definiteEnding(nounArticle Article, isPlural bool, nounCase Case) string {
	switch nounCase {
	case CaseNominative:
		if isPlural {
			return "~e"
		} else {
			switch nounArticle {
			case Der:
				return "~er"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}
		break
	case CaseAcusative:
		if isPlural {
			return "~e"
		} else {
			switch nounArticle {
			case Der:
				return "~en"
			case Die:
				return "~e"
			case Das:
				return "~es"
			}
		}
		break
	case CaseDative:
		if isPlural {
			return "~en"
		} else {
			switch nounArticle {
			case Der:
				return "~em"
			case Die:
				return "~er"
			case Das:
				return "~em"
			}
		}
		break
	case CaseGenitive:
		if isPlural {
			return "~er"
		} else {
			switch nounArticle {
			case Der:
				return "~es"
			case Die:
				return "~er"
			case Das:
				return "~es"
			}
		}
		break
	}

	return ""
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
		if isPlural {
			return "die"
		} else {
			switch nounArticle {
			case Der:
				return "der"
			case Die:
				return "die"
			case Das:
				return "das"
			}
		}
		break
	case CaseAcusative:
		if isPlural {
			return "die"
		} else {
			switch nounArticle {
			case Der:
				return "den"
			case Die:
				return "die"
			case Das:
				return "das"
			}
		}
		break
	case CaseDative:
		if isPlural {
			return "den"
		} else {
			switch nounArticle {
			case Der:
				return "dem"
			case Die:
				return "der"
			case Das:
				return "dem"
			}
		}
		break
	case CaseGenitive:
		if isPlural {
			return "der"
		} else {
			switch nounArticle {
			case Der:
				return "des"
			case Die:
				return "der"
			case Das:
				return "des"
			}
		}
		break
	}

	return ""
}
