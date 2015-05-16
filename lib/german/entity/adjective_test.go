package entity

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestAdjectiveRegexpSuccess(t *testing.T) {
	for _, testCase := range adjectiveRegexpSuccessCases {
		matches := AdjectiveRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) < 6 && matches[0] != testCase.raw {
			t.Fatalf("Regexp found: %s, expected: %s", matches[0], testCase.raw)
		}
		if matches[1] != testCase.german {
			t.Fatalf("Regexp found: %s, expected: %s", matches[1], testCase.german)
		}
		if matches[3] != testCase.comparative {
			t.Fatalf("Regexp found: %s, expected: %s", matches[3], testCase.comparative)
		}
		if matches[5] != testCase.superlative {
			t.Fatalf("Regexp found: %s, expected: %s", matches[5], testCase.superlative)
		}
	}

	t.Log(len(verbRegexpSuccessCases), "test cases")
}

func TestAdjectiveRegexpFailure(t *testing.T) {
	for _, testCase := range adjectiveRegexpFailureCases {
		matches := AdjectiveRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 {
			t.Fatalf("Regexp found: %s, expected no match", matches[0])
		}
	}

	t.Log(len(adjectiveRegexpFailureCases), "test cases")
}

func TestAdjectiveCreationSuccess(t *testing.T) {
	for num, testCase := range adjectiveCreationSuccessCases {
		adjective := NewAdjective(
			testCase.german,
			testCase.english,
			testCase.third,
			testCase.user,
			testCase.learned,
			testCase.score,
			testCase.tags,
		)

		if !reflect.DeepEqual(adjective, testCase.adjective) {
			w1, _ := json.Marshal(testCase.adjective)
			w2, _ := json.Marshal(adjective)

			t.Fatalf(
				"Adjective #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(adjectiveCreationSuccessCases), "test cases")
}

func TestAdjectiveCreationFailure(t *testing.T) {
	for num, testCase := range adjectiveCreationFailureCases {
		adjective := NewAdjective(
			testCase.german,
			testCase.english,
			testCase.third,
			testCase.user,
			testCase.learned,
			testCase.score,
			testCase.tags,
		)

		if adjective != nil {
			t.Fatalf("Adjective was created for test case number #%d.", num+1)
		}
	}
}

func TestGetComparative(t *testing.T) {
	for num, testCase := range getComparativeCases {
		result := testCase.adjective.GetComparative()

		if len(result) != len(testCase.expectedResult) {
			t.Fatalf(
				"Count of comparatives for test case #%d is different from expected. Expected: %d, got: %d",
				num+1,
				len(testCase.expectedResult),
				len(result),
			)
		}
	}

	t.Log(len(getComparativeCases), "test cases")
}

func TestGetSuperlative(t *testing.T) {
	for num, testCase := range getSuperlativeCases {
		result := testCase.adjective.GetSuperlative()

		if len(result) != len(testCase.expectedResult) {
			t.Fatalf(
				"Count of superlative for test case #%d is different from expected. Expected: %d, got: %d",
				num+1,
				len(testCase.expectedResult),
				len(result),
			)
		}
	}

	t.Log(len(getSuperlativeCases), "test cases")
}

func TestGetComparativeString(t *testing.T) {
	for num, testCase := range getComparativeCases {
		result := testCase.adjective.GetComparativeString(testCase.maxCount)

		if result != testCase.expectedResult2 {
			t.Fatalf(
				"Comparative for test case #%d is different from expected. Expected: %s, got: %s",
				num+1,
				testCase.expectedResult2,
				result,
			)
		}
	}

	t.Log(len(getComparativeCases), "test cases")
}

func TestGetSuperlativeString(t *testing.T) {
	for num, testCase := range getSuperlativeCases {
		result := testCase.adjective.GetSuperlativeString(testCase.maxCount)

		if result != testCase.expectedResult2 {
			t.Fatalf(
				"Superlative for test case #%d is different from expected. Expected: %s, got: %s",
				num+1,
				testCase.expectedResult2,
				result,
			)
		}
	}

	t.Log(len(getSuperlativeCases), "test cases")
}

func TestDeclineAdjective(t *testing.T) {
	for num, testCase := range declineAdjectiveCases {
		result := testCase.adjective.Decline(
			testCase.degree,
			testCase.declension,
			testCase.nounArticle,
			testCase.isPlural,
			testCase.nounCase,
		)

		actualResult := strings.Join(result, ", ")

		if actualResult != testCase.expectedResult {
			t.Fatalf(
				"Declension for test case #%d is different from expected. Expected: %s, got: %s",
				num+1,
				testCase.expectedResult,
				actualResult,
			)
		}
	}

	t.Log(len(declineAdjectiveCases), "test cases")
}
