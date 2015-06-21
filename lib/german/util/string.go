package util

import (
	"strings"

	generalUtil "github.com/peteraba/d5/lib/util"
)

var vowels = []string{"ie", "ei", "a", "e", "i", "o", "u", "ä", "ö", "ü", "y"}

func CountSyllables(word string) int {
	counter := 0

	for _, vowel := range vowels {
		origWord := word

		word = strings.Replace(word, vowel, "", -1)

		counter += (len(origWord) - len(word)) / len(vowel)
	}

	return counter
}

func IsVowel(char string) bool {
	for _, vowel := range vowels {
		if vowel == char {
			return true
		}
	}

	return false
}

var irregularSuffixCases = map[string][]string{
	"s":   []string{"er", "ant", "ug"},
	"n":   []string{"err", "er"},
	"nen": []string{"in"},
}

func AddSuffix(word, suffix string) string {
	if word == "" || suffix == "" {
		return word + suffix
	}

	if IsVowel(word[len(word)-1:]) || IsVowel(suffix[:1]) {
		return word + suffix
	}

	if generalUtil.HasSuffixAny(word, irregularSuffixCases[suffix]) {
		return word + suffix
	}

	return word + "e" + suffix
}
