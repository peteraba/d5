package entity

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNounRegexpSuccess(t *testing.T) {
	for _, testCase := range nounRegexpSuccessCases {
		matches := NounRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) < 3 || matches[0] != testCase.raw {
			t.Fatalf("Regexp expected: %s", testCase.raw)
		}
		if matches[1] != testCase.german {
			t.Fatalf("Regexp found: %s, expected: %s", matches[1], testCase.german)
		}
		if matches[2] != testCase.plural {
			t.Fatalf("Regexp found: %s, expected: %s", matches[2], testCase.plural)
		}
	}

	t.Log(len(nounRegexpSuccessCases), "test cases")
}

func TestNounRegexpFailure(t *testing.T) {
	for _, testCase := range nounRegexpFailureCases {
		matches := AuxiliaryRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 {
			t.Fatalf("Regexp found: %s, expected no match", matches[0])
		}
	}

	t.Log(len(nounRegexpFailureCases), "test cases")
}

func TestNounCreationSuccess(t *testing.T) {
	for num, testCase := range nounCreationSuccessCases {
		noun := NewNoun(
			testCase.articles,
			testCase.german,
			testCase.english,
			testCase.third,
			testCase.user,
			testCase.learned,
			testCase.score,
			testCase.tags,
		)

		if !reflect.DeepEqual(noun, testCase.noun) {
			w1, _ := json.Marshal(testCase.noun)
			w2, _ := json.Marshal(noun)

			t.Fatalf(
				"Noun #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(nounCreationSuccessCases), "test cases")
}

func TestGetPlural(t *testing.T) {
	for num, testCase := range getPluralCases {
		result := testCase.noun.GetPlurals()

		if len(result) != len(testCase.expectedResult) {
			t.Fatalf(
				"Count of plurals for test case #%d is different from expected. Expected: %d, got: %d",
				num+1,
				len(testCase.expectedResult),
				len(result),
			)
		}
	}

	t.Log(len(getPluralCases), "test cases")
}

func TestGetGenitive(t *testing.T) {
	for num, testCase := range getGenitiveCases {
		result := testCase.noun.GetGenitives()

		if len(result) != len(testCase.expectedResult) {
			t.Fatalf(
				"Count of genitive for test case #%d is different from expected. Expected: %d, got: %d",
				num+1,
				len(testCase.expectedResult),
				len(result),
			)
		}
	}

	t.Log(len(getGenitiveCases), "test cases")
}

func TestGetPluralString(t *testing.T) {
	for num, testCase := range getPluralCases {
		result := testCase.noun.GetPluralsString(testCase.stringCount)

		if result != testCase.expectedResult2 {
			t.Fatalf(
				"Plurals for test case #%d is different from expected. Expected: %s, got: %s",
				num+1,
				testCase.expectedResult2,
				result,
			)
		}
	}

	t.Log(len(getPluralCases), "test cases")
}

func TestGetGenitiveString(t *testing.T) {
	for num, testCase := range getGenitiveCases {
		result := testCase.noun.GetGenitivesString(testCase.stringCount)

		if result != testCase.expectedResult2 {
			t.Fatalf(
				"Genitive for test case #%d is different from expected. Expected: %s, got: %s",
				num+1,
				testCase.expectedResult2,
				result,
			)
		}
	}

	t.Log(len(getGenitiveCases), "test cases")
}
