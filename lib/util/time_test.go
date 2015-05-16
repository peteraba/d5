package util

import (
	"math"
	"testing"
	"time"
)

var parseTimeNowCases = []struct {
	timeForm, toParse string
	expectedResult    time.Time
}{
	{
		"2006-01-02",
		"2014-03-05",
		time.Date(2014, 3, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		"2006-01-02",
		"xxx",
		time.Now(),
	},
}

func TestParseTimeNow(t *testing.T) {
	for num, testCase := range parseTimeNowCases {
		actual := ParseTimeNow(testCase.timeForm, testCase.toParse).Unix()
		expected := testCase.expectedResult.Unix()

		if math.Abs(float64(actual)-float64(expected)) > 1.0 {
			t.Fatalf(
				"Parsed time #%d is different from expected.\nExpected: \n%v\ngot: \n%v",
				num+1,
				expected,
				actual,
			)
		}
	}

	t.Log(len(parseTimeNowCases), "test cases")
}
