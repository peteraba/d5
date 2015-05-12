package dict

import (
	"strings"
	"unicode/utf8"
)

var tr map[string]string

func getTr() map[string]string {
	if len(tr) != 0 {
		return tr
	}

	tr = map[string]string{
		"a": "ä",
		"o": "ö",
		"u": "ü",
	}

	return tr
}

func Decline(base, extensions string) []string {
	var (
		result = []string{}
	)

	if extensions == "" || extensions == "-" {
		return result
	}

	for _, extension := range strings.Split(extensions, "/") {
		r, size := utf8.DecodeRuneInString(extension)
		switch r {
		case '~':
			result = append(result, base+extension[size:])
			break
		case '⍨':
			result = append(result, umlautise(base)+extension[size:])
			break
		default:
			result = append(result, extension)
		}
	}

	return result
}

func umlautise(base string) string {
	var (
		tr       = getTr()
		fromIdx  int
		foundIdx int    = -1
		foundKey string = ""
	)

	for from, _ := range tr {
		fromIdx = strings.LastIndex(base, from)
		if fromIdx > foundIdx {
			foundIdx = fromIdx
			foundKey = from
		}
	}

	if foundIdx > -1 {
		base = strings.Replace(base, foundKey, tr[foundKey], 1)
	}

	return base
}
