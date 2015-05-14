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
