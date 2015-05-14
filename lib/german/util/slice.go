package util

import "strings"

func SliceAppend(stringSlice []string, stringToAppend string) []string {
	result := []string{}
	word := ""

	for _, stem := range stringSlice {
		word = stem + stringToAppend
		switch stringToAppend {
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
