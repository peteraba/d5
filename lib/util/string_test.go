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
			t.Fatalf(
				"Case #%d: Number of strings split is wrong. Expected %d, got %d.",
				num,
				len(testCase.result),
				len(result),
			)
		}

		for num, word := range result {
			if word != testCase.result[num] {
				t.Fatalf(
					"Case #%d: Regexp found: %s, expected: %s",
					num,
					word,
					testCase.result[num],
				)
			}
		}
	}

	t.Log(len(trimSplitCases), "test cases")
}

var joinLimitedCases = []struct {
	parts    []string
	maxCount int
	result   string
}{
	{
		[]string{"hello"},
		0,
		"",
	},
	{
		[]string{"hello"},
		2,
		"hello",
	},
	{
		[]string{"hello", "hi"},
		2,
		"hello, hi",
	},
	{
		[]string{"hello", "hi", "ola"},
		2,
		"hello, hi",
	},
}

func TestJoinLimited(t *testing.T) {
	for num, testCase := range joinLimitedCases {
		result := JoinLimited(testCase.parts, ", ", testCase.maxCount)

		if result != testCase.result {
			t.Fatalf("Case #%d. Expected: '%v', got: '%v'\n", num, testCase.result, result)
		}
	}

	t.Log(len(joinLimitedCases), "test cases")
}
