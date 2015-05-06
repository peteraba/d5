package german

import (
	"encoding/json"
	"reflect"
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
