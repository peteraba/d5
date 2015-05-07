package entity

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMeaningRegexpSuccess(t *testing.T) {
	for _, testCase := range meaningRegexpSuccessCases {
		matches := MeaningRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) == 0 {
			t.Fatalf("Regexp match failed. expected: '%s'", testCase.raw)
		}
		if len(matches) < 4 || matches[0] != testCase.raw {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[0], testCase.raw)
		}
		if matches[1] != testCase.main {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[1], testCase.main)
		}
		if matches[3] != testCase.parant {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[3], testCase.parant)
		}
	}

	t.Log(len(meaningRegexpSuccessCases), "test cases")
}

func TestMeaningRegexpFailure(t *testing.T) {
	for _, testCase := range meaningRegexpFailureCases {
		matches := MeaningRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 && matches[0] != "" {
			t.Fatalf("Regexp found: '%s', expected no match", matches[0])
		}
	}

	t.Log(len(meaningRegexpFailureCases), "test cases")
}

func TestMeaningCreationSuccess(t *testing.T) {
	for num, testCase := range meaningCreationCases {
		meanings := NewMeanings(testCase.allMeanings)

		if !reflect.DeepEqual(meanings, testCase.meanings) {
			w1, _ := json.Marshal(testCase.meanings)
			w2, _ := json.Marshal(meanings)

			t.Fatalf(
				"Meaning list #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(meaningRegexpSuccessCases), "test cases")
}

func TestWordCreationSuccess(t *testing.T) {
	for num, testCase := range wordCreationSuccessCases {
		word := NewAny(
			testCase.german,
			testCase.english,
			testCase.third,
			testCase.category,
			testCase.user,
			testCase.learned,
			testCase.score,
			testCase.tags,
			testCase.errors,
		)

		if !reflect.DeepEqual(word, testCase.word) {
			w1, _ := json.Marshal(testCase.word)
			w2, _ := json.Marshal(word)

			t.Fatalf(
				"Word #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(wordCreationSuccessCases), "test cases")
}