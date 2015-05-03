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

func TestAuxiliaryRegexp(t *testing.T) {
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

func TestNounRegexp(t *testing.T) {
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

func TestAdjectiveRegexp(t *testing.T) {
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

func TestVerbRegexp(t *testing.T) {
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
