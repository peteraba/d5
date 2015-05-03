package util

import "testing"

var trimSplitCases = []struct {
	input     string
	result    []string
	separator string
}{
	{
		"",
		[]string{},
		"",
	},
	{
		"    hello   |  hi   ",
		[]string{"hello", "hi"},
		"|",
	},
	{
		"    hello   |  hi  | yo! ",
		[]string{"hello", "hi", "yo!"},
		"|",
	},
}

func TestTripSplit(t *testing.T) {
	for num, testCase := range trimSplitCases {
		result := TrimSplit(testCase.input, testCase.separator)

		if len(result) != len(testCase.result) {
			t.Fatalf("Case #%d: Number of strings split is wrong. Expected %d, got %d.", num, len(testCase.result), len(result))
		}

		for num, word := range result {
			if word != testCase.result[num] {
				t.Fatalf("Case #%d: Regexp found: %s, expected: %s", num, word, testCase.result[num])
			}
		}
	}

	t.Log(len(trimSplitCases), "test cases")
}
