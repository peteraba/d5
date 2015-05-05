package german

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestArticleRegexpSucecss(t *testing.T) {
	for _, testCase := range articleRegexpSuccessCases {
		matches := ArticleRegexp.FindStringSubmatch(testCase.raw)
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

	t.Log(len(articleRegexpSuccessCases), "test cases")
}

func TestArticleRegexpFailure(t *testing.T) {
	for _, testCase := range articleRegexpFailureCases {
		matches := ArticleRegexp.FindStringSubmatch(testCase)
		if len(matches) > 0 {
			t.Fatalf("Regexp found: %s, expected no match", matches[0])
		}
	}

	t.Log(len(articleRegexpFailureCases), "test cases")
}

func TestAuxiliaryRegexpSuccess(t *testing.T) {
	for _, testCase := range auxiliaryRegexpSuccessCases {
		matches := AuxiliaryRegexp.FindStringSubmatch(testCase.raw)
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

func TestVerbRegexpSuccess(t *testing.T) {
	for _, testCase := range verbRegexpSuccessCases {
		matches := VerbRegexp.FindStringSubmatch(testCase.raw)
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
		arguments := NewArgument(testCase.allArguments)

		if !reflect.DeepEqual(arguments, testCase.arguments) {
			w1, _ := json.Marshal(arguments)
			w2, _ := json.Marshal(testCase.arguments)

			t.Fatalf(
				"Argument list #%d is different from expected.\nExpected: %v\ngot: %v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(meaningRegexpSuccessCases), "test cases")
}

func TestMeaningCreationSuccess(t *testing.T) {
	for num, testCase := range meaningCreationCases {
		meanings := NewMeanings(testCase.allMeanings)

		if !reflect.DeepEqual(meanings, testCase.meanings) {
			w1, _ := json.Marshal(meanings)
			w2, _ := json.Marshal(testCase.meanings)

			t.Fatalf(
				"Meaning list #%d is different from expected.\nExpected: %v\ngot: %v",
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
			w1, _ := json.Marshal(word)
			w2, _ := json.Marshal(testCase.word)

			t.Fatalf(
				"Word #%d is different from expected.\nExpected: %v\ngot: %v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(wordCreationSuccessCases), "test cases")
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
			w1, _ := json.Marshal(adjective)
			w2, _ := json.Marshal(testCase.adjective)

			t.Fatalf(
				"Adjective #%d is different from expected.\nExpected: %v\ngot: %v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(adjectiveCreationSuccessCases), "test cases")
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
			w1, _ := json.Marshal(noun)
			w2, _ := json.Marshal(testCase.noun)

			t.Fatalf(
				"Noun #%d is different from expected.\nExpected: %v\ngot: %v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(nounCreationSuccessCases), "test cases")
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

		if verb == nil {
			t.Fatalf("No verb is created for case #%d, german word: %s", num, testCase.german)
		}
	}

	t.Log(len(verbCreationSuccessCases), "test cases")
}
