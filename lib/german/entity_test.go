package german

import "testing"

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
	var testArgument Argument

	for num, testCase := range argumentCreationCases {
		arguments := NewArgument(testCase.allArguments)

		if len(arguments) != len(testCase.arguments) {
			t.Fatal(arguments)
			t.Fatalf(
				"Case #%d. Wrong number of arguments created. Expected %d, got: %d.",
				num,
				len(testCase.arguments),
				len(arguments),
			)
		}

		for key, argument := range arguments {
			testArgument = testCase.arguments[key]

			if argument.Preposition != testArgument.Preposition {
				t.Fatalf(
					"Case #%d. Wrong preposition found. Expected '%s', got: '%s'",
					num,
					testArgument.Preposition,
					argument.Case,
				)
			}

			if argument.Case != testArgument.Case {
				t.Fatalf(
					"Case #%d. Wrong case found. Expected '%s', got: '%s'",
					num,
					testArgument.Case,
					argument.Case,
				)
			}
		}
	}

	t.Log(len(meaningRegexpSuccessCases), "test cases")
}

func TestMeaningCreationSuccess(t *testing.T) {
	var testMeaning Meaning

	for num, testCase := range meaningCreationCases {
		meanings := NewMeanings(testCase.allMeanings)

		if len(meanings) != len(testCase.meanings) {
			t.Fatalf("Case #%d: Wrong number of meanings created. Expected %d, got: %d.", num, testCase.meanings, meanings)
		}

		for key, meaning := range meanings {
			testMeaning = testCase.meanings[key]

			if meaning.Main != testMeaning.Main {
				t.Fatalf(
					"Case #%d. Wrong main found. Expected '%s', got: '%s'",
					num,
					testMeaning.Main,
					meaning.Main,
				)
			}

			if meaning.Parantheses != testMeaning.Parantheses {
				t.Fatalf(
					"Case #%d. Wrong parantheses found. Expected '%s', got: '%s'",
					num,
					testMeaning.Parantheses,
					meaning.Parantheses,
				)
			}
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

		if word == nil {
			t.Fatalf("No word is created for case #%d, german word: %s", num, testCase.german)
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

		if adjective == nil {
			t.Fatalf("No adjective is created for case #%d, german word: %s", num, testCase.german)
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

		if noun == nil {
			t.Fatalf("No noun is created for case #%d, german word: %s", num, testCase.german)
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
