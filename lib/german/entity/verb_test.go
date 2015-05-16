package entity

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	germanUtil "github.com/peteraba/d5/lib/german/util"
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

func TestPrefixCreation(t *testing.T) {
	for num, testCase := range prefixCreationCases {
		prefix := NewPrefix(testCase.german)

		if prefix.Prefix != testCase.prefix.Prefix && prefix.Separable != testCase.prefix.Separable {
			t.Fatalf(
				"Prefix #%d isn't as expected. Expected: '%v', got: '%v'.",
				num+1,
				testCase.prefix,
				prefix,
			)
		}
	}

	t.Log(len(prefixCreationCases), "test cases")
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

func TestVerbCreationFailure(t *testing.T) {
	for num, testCase := range verbCreationFailureCases {
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

		if verb != nil {
			t.Fatalf(
				"Verb #%d is not nil as expected.\nGot: \n%v",
				num+1,
				verb,
			)
		}
	}

	t.Log(len(verbRegexpFailureCases), "test cases")
}

func TestConjugation(t *testing.T) {
	for _, testCase := range verbCreationSuccessCases {
		verb := testCase.verb

		conjugationCheck(t, verb.German, testCase.presentS1, verb.GetPresentS1(), "Present S1")
		conjugationCheck(t, verb.German, testCase.presentS2, verb.GetPresentS2(), "Present S2")
		conjugationCheck(t, verb.German, testCase.presentS3, verb.GetPresentS3(), "Present S3")
		conjugationCheck(t, verb.German, testCase.presentP1, verb.GetPresentP1(), "Present P1")
		conjugationCheck(t, verb.German, testCase.presentP2, verb.GetPresentP2(), "Present P2")
		conjugationCheck(t, verb.German, testCase.presentP3, verb.GetPresentP3(), "Present P3")

		conjugationCheck(t, verb.German, testCase.preteriteS1, verb.GetPreteriteS1(), "Preterite S1")
		conjugationCheck(t, verb.German, testCase.preteriteS2, verb.GetPreteriteS2(), "Preterite S2")
		conjugationCheck(t, verb.German, testCase.preteriteS3, verb.GetPreteriteS3(), "Preterite S3")
		conjugationCheck(t, verb.German, testCase.preteriteP1, verb.GetPreteriteP1(), "Preterite P1")
		conjugationCheck(t, verb.German, testCase.preteriteP2, verb.GetPreteriteP2(), "Preterite P2")
		conjugationCheck(t, verb.German, testCase.preteriteP3, verb.GetPreteriteP3(), "Preterite P3")
	}

	t.Log(len(verbCreationSuccessCases), "test cases")
}

func conjugationCheck(t *testing.T, german string, expected []string, actual []string, str string) {
	var stringsExpected = strings.Join(expected, ",")
	var stringsActual = strings.Join(actual, ",")

	if stringsExpected != stringsActual {
		t.Fatalf(
			"%s for %s is wrong. Expected: '%s', got: '%s'.",
			str,
			german,
			stringsExpected,
			stringsActual,
		)
	}
}

func TestConjugationSeparated(t *testing.T) {
	var (
		actual [][2]string
	)

	for _, testCase := range verbCreationSuccessCases {
		verb := testCase.verb

		actual = testCase.verb.GetSeparated(testCase.pp, testCase.tense)

		conjugationSeparatedCheck(t, verb.German, testCase.expectedSeparated, actual)
	}

	t.Log(len(verbCreationSuccessCases), "test cases")
}

func conjugationSeparatedCheck(t *testing.T, german string, expected [][2]string, actual [][2]string) {
	var (
		stringExpected string
		stringActual   string
	)

	stringExpected = germanUtil.JoinSeparatedList(expected, "|", 1, ", ")
	stringActual = germanUtil.JoinSeparatedList(actual, "|", 1, ", ")

	if stringExpected != stringActual {
		t.Fatalf(
			"Separated word for '%s' is wrong. Expected: '%s', got: '%s'.",
			german,
			stringExpected,
			stringActual,
		)
	}
}
