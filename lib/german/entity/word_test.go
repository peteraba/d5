package entity

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/peteraba/d5/lib/general"
)

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

func TestMeaningCreationSuccess(t *testing.T) {
	for num, testCase := range meaningCreationSuccessCases {
		meanings, _ := NewMeanings(testCase.allMeanings, []string{})

		if !reflect.DeepEqual(meanings, testCase.meanings) {
			w1, _ := json.Marshal(testCase.meanings)
			w2, _ := json.Marshal(meanings)

			t.Fatalf(
				"Meaning list #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(meaningCreationSuccessCases), "test cases")
}

func TestMeaningCreationFailure(t *testing.T) {
	for num, testCase := range meaningCreationFailureCases {
		_, errors := NewMeanings(testCase.allMeanings, []string{})

		if len(errors) == 0 {
			w1, _ := json.Marshal(testCase.allMeanings)

			t.Fatalf(
				"Meaning list #%d did not return errors as expected.\nMeaning: \n%v",
				num+1,
				string(w1),
			)
		}
	}

	t.Log(len(meaningCreationFailureCases), "test cases")
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
			w1, _ := json.Marshal(testCase.word)
			w2, _ := json.Marshal(word)

			t.Fatalf(
				"Word #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				string(w1),
				string(w2),
			)
		}
	}

	t.Log(len(wordCreationSuccessCases), "test cases")
}

func TestWordGetEnglish(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetEnglish()
		if !reflect.DeepEqual(actual, word.English) {
			t.Fatal(
				"GetEnglish test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.English,
				actual,
			)
		}
	}
}

func TestWordGetThird(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetThird()
		if !reflect.DeepEqual(actual, word.Third) {
			t.Fatal(
				"GetThird test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.Third,
				actual,
			)
		}
	}
}

func TestWordGetCategory(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetCategory()
		if actual != word.Category {
			t.Fatal(
				"GetCategory test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.Category,
				actual,
			)
		}
	}
}

func TestWordGetScore(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetScore()
		if actual != word.Score {
			t.Fatal(
				"GetScore test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.Score,
				actual,
			)
		}
	}
}

func TestWordGetUser(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetUser()
		if actual != word.User {
			t.Fatal(
				"GetUser test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.User,
				actual,
			)
		}
	}
}

func TestWordGetLearned(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetLearned()
		if actual != word.Learned {
			t.Fatal(
				"GetGerman test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.Learned,
				actual,
			)
		}
	}
}

func TestWordGetErrors(t *testing.T) {
	var word *Any

	for num, testCase := range wordCreationSuccessCases {
		word = testCase.word

		actual := word.GetErrors()
		if !reflect.DeepEqual(actual, word.Errors) {
			t.Fatalf(
				"GetErrors test failed for case #%d. Expected: '%v', got: '%v'.",
				num+1,
				word.Errors,
				actual,
			)
		}
	}
}

func TestAddScoreAddsScore(t *testing.T) {
	var (
		sut   = DefaultWord{}
		score = general.Score{}
	)

	score.Result = 6
	score.LearnedAt = time.Now()

	sut.AddScore(&score)

	if len(sut.GetScores()) != 1 {
		t.Fatalf("Adding score failed. Expected to have 1 score, %d found.", len(sut.GetScores()))
	}
}
