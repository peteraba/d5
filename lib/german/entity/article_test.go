package entity

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

func TestIndefiniteArticle(t *testing.T) {
	for num, testCase := range indefiniteArticleCases {
		article := IndefiniteArticle(testCase.word, testCase.nounArticle, testCase.isPlural, testCase.nounCase)

		if article != testCase.result {
			t.Fatalf(
				"Article found is wrong for case #%d. Expected: '%s', got: '%s'",
				num+1,
				testCase.result,
				article,
			)
		}
	}

	t.Log(len(indefiniteArticleCases), "test cases")
}

func TestDefiniteArticle(t *testing.T) {
	for num, testCase := range definiteArticleCases {
		article := DefiniteArticle(testCase.nounArticle, testCase.isPlural, testCase.nounCase)

		if article != testCase.result {
			t.Fatalf(
				"Article found is wrong for case #%d. Expected: '%s', got: '%s'",
				num+1,
				testCase.result,
				article,
			)
		}
	}

	t.Log(len(definiteArticleCases), "test cases")
}
