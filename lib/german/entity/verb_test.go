package entity

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAuxiliaryRegexpSuccess(t *testing.T) {
	for _, testCase := range auxiliaryRegexpSuccessCases {
		matches := AuxiliaryRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) == 0 {
			t.Fatalf("Regexp found: no match")
		}
		if len(matches) < 4 || matches[0] != testCase.raw {
			t.Fatalf("Regexp found: %s, expected: %s", matches[0], testCase.raw)
		}
		if matches[1] != testCase.first {
			t.Fatalf("Regexp found: %s, expected: %s", matches[1], testCase.first)
		}
		if matches[3] != testCase.second {
			t.Fatalf("Regexp found: %s, expected: %s", matches[3], testCase.second)
		}
	}

	t.Log(len(auxiliaryRegexpSuccessCases), "test cases")
}

func TestAuxiliaryRegexpFailure(t *testing.T) {
	for _, testCase := range auxiliaryRegexpFailureCases {
		matches := AuxiliaryRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 {
			t.Fatalf("Regexp found: %s, expected no match", matches[0])
		}
	}

	t.Log(len(auxiliaryRegexpFailureCases), "test cases")
}

func TestArgumentRegexpSuccess(t *testing.T) {
	for _, testCase := range argumentRegexpSuccessCases {
		matches := ArgumentRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) == 0 {
			t.Fatalf("Regexp found: no match")
		}
		if len(matches) < 4 || matches[0] != testCase.raw {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[0], testCase.raw)
		}
		if matches[1] != testCase.argPrep {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[1], testCase.argPrep)
		}
		if matches[3] != testCase.argCase {
			t.Fatalf("Regexp found: '%s', expected: '%s'", matches[3], testCase.argCase)
		}
	}

	t.Log(len(argumentRegexpSuccessCases), "test cases")
}

func TestArgumentRegexpFailure(t *testing.T) {
	for _, testCase := range argumentRegexpFailureCases {
		matches := ArgumentRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 && matches[0] != "" {
			t.Fatalf("Regexp found: '%s', expected no match", matches[0])
		}
	}

	t.Log(len(argumentRegexpFailureCases), "test cases")
}

func TestVerbRegexpSuccess(t *testing.T) {
	for _, testCase := range verbRegexpSuccessCases {
		matches := VerbRegexp.FindStringSubmatch(testCase.raw)
		if len(matches) == 0 {
			t.Fatalf("Regexp found: no match")
		}
		if matches[0] != testCase.raw {
			t.Fatalf("Regexp found: %s, expected: %s", matches[0], testCase.raw)
		}
		if matches[1] != testCase.german {
			t.Fatalf("Regexp found: %s, expected: %s", matches[1], testCase.german)
		}
		if matches[2] != testCase.arguments {
			t.Fatalf("Regexp found: %s, expected: %s", matches[2], testCase.arguments)
		}
	}

	t.Log(len(verbRegexpSuccessCases), "test cases")
}

func TestVerbRegexpFailure(t *testing.T) {
	for _, testCase := range verbRegexpFailureCases {
		matches := VerbRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 {
			t.Fatalf("Regexp found: %s, expected no match", matches[0])
		}
	}

	t.Log(len(verbRegexpFailureCases), "test cases")
}

func TestArgumentCreationSuccess(t *testing.T) {
	for num, testCase := range argumentCreationCases {
		arguments := NewArguments(testCase.allArguments)

		if !reflect.DeepEqual(arguments, testCase.arguments) {
			w1, _ := json.Marshal(testCase.arguments)
			w2, _ := json.Marshal(arguments)

			t.Fatalf(
				"Argument list #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(meaningRegexpSuccessCases), "test cases")
}

func TestVerbCreationSuccess(t *testing.T) {
	for num, testCase := range verbCreationSuccessCases {
		verb := NewVerb(
			testCase.auxiliary,
			testCase.german,
			testCase.english,
			testCase.third,
			testCase.user,
			testCase.learned,
			testCase.score,
			testCase.tags,
		)

		if !reflect.DeepEqual(verb, testCase.verb) {
			w1, _ := json.Marshal(testCase.verb)
			w2, _ := json.Marshal(verb)

			t.Fatalf(
				"Verb #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(verbCreationSuccessCases), "test cases")
}
