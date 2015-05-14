package dict

import "testing"

var declensionCases = []struct {
	base      string
	extension string
	result    string
}{
	{
		"jung",
		"⍨er",
		"jünger",
	},
	{
		"schmal",
		"~er",
		"schmaler",
	},
	{
		"zusätzlich",
		"-",
		"",
	},
	{
		"Stipendium",
		"Stipendien",
		"Stipendien",
	},
	{
		"sauber",
		"⍨er",
		"saüberer",
	},
}

func TestDecline(t *testing.T) {
	for num, testCase := range declensionCases {
		result := Decline(testCase.base, testCase.extension)

		if result != testCase.result {
			t.Fatalf("Case %d: result is different from expected. Expected: %s result(s), got: %s.", num, testCase.result, result)
		}
	}

	t.Log(len(declensionCases), "test cases")
}
