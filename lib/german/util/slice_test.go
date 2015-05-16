package util

import (
	"strings"
	"testing"
)

var sliceAppendCases = []struct {
	stringSlice    []string
	stringToAppend string
	expectedResult string
}{
	{
		[]string{"aufbau"},
		"t",
		"aufbaut",
	},
	{
		[]string{"ausarbeit"},
		"t",
		"ausarbeitet",
	},
	{
		[]string{"arbeite"},
		"en",
		"arbeiten",
	},
}

func TestSliceAppend(t *testing.T) {
	var (
		actualResult []string
	)

	for num, testCase := range sliceAppendCases {
		actualResult = SliceAppend(testCase.stringSlice, testCase.stringToAppend)

		if strings.Join(actualResult, ", ") != testCase.expectedResult {
			t.Fatalf(
				"Failed to append slice properly, test case #%d. Expected: '%s', got: '%s'.",
				num,
				testCase.expectedResult,
				actualResult,
			)
		}
	}

	t.Log(len(sliceAppendCases), "test cases")
}

var joinSeparatedListCases = []struct {
	wordList        [][2]string
	joinSeparatedBy string
	first           int
	joinListBy      string
	expectedResult  string
}{
	{
		[][2]string{
			[2]string{"arbeit", "aus"},
			[2]string{"bau", "auf"},
		},
		"|",
		3,
		", ",
		"arbeit|aus, bau|auf",
	},
	{
		[][2]string{
			[2]string{"arbeit", "aus"},
			[2]string{"bau", "auf"},
		},
		"|",
		1,
		", ",
		"aus|arbeit, auf|bau",
	},
}

func TestJoinSeparatedList(t *testing.T) {
	var actualResult string

	for num, testCase := range joinSeparatedListCases {
		actualResult = JoinSeparatedList(
			testCase.wordList,
			testCase.joinSeparatedBy,
			testCase.first,
			testCase.joinListBy,
		)

		if actualResult != testCase.expectedResult {
			t.Fatalf(
				"Failed to join separated list properly, test case #%d. Expected: '%s', got: '%s'.",
				num,
				testCase.expectedResult,
				actualResult,
			)
		}
	}

	t.Log(len(joinSeparatedListCases), "test cases")
}
