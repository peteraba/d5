package dict

import "testing"

var declensionCases = []struct {
	base      string
	extension string
	result    []string
}{
	{
		"jung",
		"⍨er",
		[]string{"jünger"},
	},
	{
		"schmal",
		"~er/⍨er",
		[]string{"schmaler", "schmäler"},
	},
	{
		"zusätzlich",
		"-",
		[]string{},
	},
	{
		"Stipendium",
		"Stipendien",
		[]string{"Stipendien"},
	},
}

func TestExtend(t *testing.T) {
	for num, testCase := range declensionCases {
		result := Decline(testCase.base, testCase.extension)

		if len(result) != len(testCase.result) {
			t.Fatalf("Case %d: result is different from expected. Expected: %d result(s), got: %d.", num, len(testCase.result), len(result))
		}

		for key, word := range result {
			if word != testCase.result[key] {
				t.Fatalf("Case %d / %d. Expected: '%s', got: '%s'", num, key, testCase.result[key], word)
			}
		}
	}

	t.Log(len(declensionCases), "test cases")
}
