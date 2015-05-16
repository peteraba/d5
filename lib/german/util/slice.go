package util

import "strings"

func SliceAppend(stringSlice []string, stringToAppend string) []string {
	result := []string{}
	word := ""

	for _, stem := range stringSlice {
		word = stem + stringToAppend
		switch stringToAppend {
		case "en":
			if strings.HasSuffix(stem, "e") {
				word = strings.TrimSuffix(stem, "e") + stringToAppend
			}
			break
		case "t":
			if strings.HasSuffix(stem, "t") {
				word = stem + "e" + stringToAppend
			}
			break
		}
		result = append(result, word)
	}

	return result
}

func JoinSeparated(words [2]string, joinBy string, first int) string {
	if first < 0 || first > 1 {
		first = 0
	}

	return words[first] + joinBy + words[1-first]
}

func JoinSeparatedList(wordsList [][2]string, joinSeparatedBy string, first int, joinListBy string) string {
	var result = []string{}

	for _, word := range wordsList {
		result = append(result, JoinSeparated(word, joinSeparatedBy, first))
	}

	return strings.Join(result, joinListBy)
}
