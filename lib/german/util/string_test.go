package util

import "testing"

var countSyllablesCases = []struct {
	word          string
	expectedCount int
}{
	{"fünf", 1},
	{"hallo", 2},
	{"Bäckerei", 3},
}

func TestPluralGenitiveDeclension(t *testing.T) {
	for num, testCase := range countSyllablesCases {
		result := CountSyllables(testCase.word)

		if testCase.expectedCount != result {
			t.Fatalf(
				"Counting syllables of test case #%d is different from the expected. Expected: %d, got: %d.",
				num+1,
				testCase.expectedCount,
				result,
			)
		}
	}

	t.Log(len(countSyllablesCases), "test cases")
}

func TestIsVowel(t *testing.T) {
	if IsVowel("b") {
		t.Fatal("B was found to be a vowel")
	}

	t.Log(1, "test case")
}

func TestIsVowelReturnsFalseForConsonants(t *testing.T) {
	if !IsVowel("a") {
		t.Fatal("A was found to be a vowel")
	}

	t.Log(1, "test case")
}

var addSuffixCases = []struct {
	word           string
	suffix         string
	expectedResult string
}{
	{"", "suf", "suf"},
	{"Abo", "s", "Abos"},
	{"Abfahrt", "n", "Abfahrten"},
	{"Herr", "n", "Herrn"},
	{"Ägypter", "s", "Ägypters"},
	{"Ägypter", "n", "Ägyptern"},
	{"Croissant", "s", "Croissants"},
	{"Brasilianerin", "nen", "Brasilianerinnen"},
	{"Umsatz", "s", "Umsatzes"},
	{"", "", ""},
}

func TestAddSuffix(t *testing.T) {
	for num, testCase := range addSuffixCases {
		actualResult := AddSuffix(testCase.word, testCase.suffix)

		if actualResult != testCase.expectedResult {
			t.Fatalf(
				"Adding suffix in test case #%d failed. Expected: '%s', got '%s'",
				num+1,
				testCase.expectedResult,
				actualResult,
			)
		}
	}

	t.Log(len(addSuffixCases), "test case")
}
