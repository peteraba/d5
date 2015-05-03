package german

import "strings"

func TrimSplit(s, sep string) []string {
	split := strings.Split(s, sep)

	for key, word := range split {
		split[key] = strings.Trim(word, " \n\t")
	}

	return split
}
