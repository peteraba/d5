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

func TestMeaningCreationSuccess(t *testing.T) {
	//for _, testCase := range meaningRegexpSuccessCases {

	//}

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
			testCase.ok,
		)

		if word == nil {
			t.Fatalf("No word is created for case #%d, german word: %s", num, testCase.german)
		}
	}

	t.Log(len(wordCreationSuccessCases), "test cases")
}
