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

func Decline(base, extension string) string {
	if extension == "" || extension == "-" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(extension)
	switch r {
	case '~':
		return base + extension[size:]
		break
	case '⍨':
		return umlautise(base) + extension[size:]
		break
	}

	return extension
}

func umlautise(base string) string {
	var (
		tr              = getTr()
		fromIdx  int    = 0
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
