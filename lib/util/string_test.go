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

var hasSuffixAnyCases = []struct {
	s              string
	suffixes       []string
	expectedResult bool
}{
	{
		"hello",
		[]string{"ola", "lo"},
		true,
	},
	{
		"hello",
		[]string{"ye", "ma"},
		false,
	},
}

func TestHasSuffix(t *testing.T) {
	for num, testCase := range hasSuffixAnyCases {
		actualResult := HasSuffixAny(testCase.s, testCase.suffixes)

		if actualResult != testCase.expectedResult {
			t.Fatalf("Case #%d failed. Expected: '%v', got: '%v'\n", num, testCase.expectedResult, actualResult)
		}
	}

	t.Log(len(hasSuffixAnyCases), "test cases")
}

var stringInCases = []struct {
	s              string
	options        []string
	expectedResult bool
}{
	{
		"hello",
		[]string{"nope", "hello", "yeah"},
		true,
	},
	{
		"hello",
		[]string{"hell", "ello", "hello!"},
		false,
	},
}

func TestStringIn(t *testing.T) {
	for num, testCase := range stringInCases {
		actualResult := StringIn(testCase.s, testCase.options)

		if actualResult != testCase.expectedResult {
			t.Fatalf("Case #%d failed. Expected: '%v', got: '%v'\n", num, testCase.expectedResult, actualResult)
		}
	}

	t.Log(len(stringInCases), "test cases")
}
