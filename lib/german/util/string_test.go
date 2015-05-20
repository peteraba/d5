package util

import "testing"

var countSyllablesCases = []struct {
	word          string
	expectedCount int
}{
	{"hello", 2},
}

func TestPluralGenitiveDeclension(t *testing.T) {
	for num, testCase := range countSyllablesCases {
		result := CountSyllables(testCase.word)

		if testCase.expectedCount != result {
			t.Fatalf(
				"Counting syllables of test case #%d is not expected. Expected: %d, got: %d.",
				num+1,
				testCase.expectedCount,
				result,
			)
		}
	}

	t.Log(len(countSyllablesCases), "test cases")
}
