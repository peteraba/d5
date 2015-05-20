package entity

import (
	"encoding/json"
	"reflect"
	"strings"
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

func TestNounCreationFailure(t *testing.T) {
	var actual = NewNoun("", "", "", "", "", "", "", "")

	if actual != nil {
		t.Fatal("Noun should have not been created with empty German argument")
	}
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

func TestSingularNominativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(false, CaseNominative)
		expected := strings.Join(testCase.singularNominative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Singular nominative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestSingularAcusativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(false, CaseAcusative)
		expected := strings.Join(testCase.singularAcusative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Singular acusative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestSingularDativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(false, CaseDative)
		expected := strings.Join(testCase.singularDative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Singular dative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestSingularGenitiveDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(false, CaseGenitive)
		expected := strings.Join(testCase.singularGenitive, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Singular genitive declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestPluralNominativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(true, CaseNominative)
		expected := strings.Join(testCase.pluralNominative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Plural nominative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestPluralAcusativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(true, CaseAcusative)
		expected := strings.Join(testCase.pluralAcusative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Plural acusative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestPluralDativeDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(true, CaseDative)
		expected := strings.Join(testCase.pluralDative, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Plural dative declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}

func TestPluralGenitiveDeclension(t *testing.T) {
	for num, testCase := range getNounDeclensionCases {
		result := testCase.noun.Decline(true, CaseGenitive)
		expected := strings.Join(testCase.pluralGenitive, ",")
		actual := strings.Join(result, ",")

		if expected != actual {
			t.Fatalf(
				"Plural genitive declension of test case #%d is not as expected. Expected: '%s', got: '%s'.",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(getNounDeclensionCases), "test cases")
}
