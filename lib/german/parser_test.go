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
	}

	t.Log(len(parseWordCases), "test cases")
}
