package util

import "regexp"

const vowels = "aeiouäöü"

var regexpVowels = regexp.MustCompile("[aeiouäöü]")

func CountSyllables(word string) int {
	noVowels := regexpVowels.ReplaceAllString(word, "")

	return len(word) - len(noVowels)
}
