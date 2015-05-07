package entity

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
