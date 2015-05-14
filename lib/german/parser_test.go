package german

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/peteraba/d5/lib/german/entity"
)

func newEmptySuperword(category string) superword {
	return superword{
		entity.DefaultWord{
			"",
			[]entity.Meaning{},
			[]entity.Meaning{},
			category,
			"",
			time.Now(),
			5,
			[]string{},
			[]string{},
		},
		[]entity.Auxiliary{},
		entity.Prefix{},
		"",
		"",
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		entity.ReflexiveWithout,
		[]entity.Argument{},
		[]entity.Article{},
		[]string{},
		[]string{},
		false,
		[]string{},
		[]string{},
	}
}

func newEmptyIdiom() entity.DefaultWord {
	w := entity.DefaultWord{}

	w.Category = "idiom"

	return w
}

func newEmptyNoun() entity.Noun {
	return entity.Noun{}
}

func newEmptyAdjective() entity.Adjective {
	return entity.Adjective{}
}

func newEmptyVerb() entity.Verb {
	return entity.Verb{}
}

var parseWordCases = []struct {
	superwords []superword
	words      []entity.Word
}{
	{
		[]superword{
			newEmptySuperword("verb"),
		},
		[]entity.Word{
			newEmptyVerb(),
		},
	},
	{
		[]superword{
			newEmptySuperword("idiom"),
			newEmptySuperword("noun"),
			newEmptySuperword("adj"),
			newEmptySuperword("idiom"),
		},
		[]entity.Word{
			newEmptyIdiom(),
			newEmptyNoun(),
			newEmptyAdjective(),
			newEmptyIdiom(),
		},
	},
}

func TestSliceAppend(t *testing.T) {
	for num, testCase := range parseWordCases {
		b, err := json.Marshal(testCase.superwords)

		if err != nil {
			t.Fatal(err)
		}

		words, err := ParseWords(b)

		if err != nil {
			t.Fatal(err)
		}

		if len(words) != len(testCase.words) {
			t.Fatalf(
				"Test case #%d is wrong. Expected %d, got: %d",
				num,
				len(testCase.words),
				len(words),
			)
		}

		for num2, w := range words {
			if w.GetGerman() != testCase.words[num2].GetGerman() {
				t.Fatalf(
					"Word #%d, test case #%d is wrong. Expected %s, got: %s",
					num2,
					num,
					testCase.words[num2],
					w,
				)
			}
		}
	}

	t.Log(len(parseWordCases), "test cases")
}

func TestSliceAppendErrors(t *testing.T) {
	if err, _ := ParseWords([]byte{}); err == nil {
		t.Fatal("Empty byte slice should cause an error")
	}
}
