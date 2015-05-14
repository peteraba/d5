package entity

var articleRegexpSuccessCases = []struct {
	raw, first, second string
}{
	{"s", "s", ""},
	{"s/r", "s", "r"},
	{"r/s", "r", "s"},
}

var articleRegexpFailureCases = []string{
	"",
	"i",
	"s/r/s",
	"S",
}

var indefiniteArticleCases = []struct {
	word        string
	nounArticle Article
	isPlural    bool
	nounCase    Case
	result      string
}{
	{
		"ein",
		Der,
		false,
		CaseNominative,
		"ein",
	},
	{
		"ein",
		Der,
		false,
		CaseAcusative,
		"einen",
	},
	{
		"ein",
		Der,
		true,
		CaseDative,
		"",
	},
	{
		"ein",
		Die,
		false,
		CaseDative,
		"einer",
	},
	{
		"ein",
		Das,
		false,
		CaseDative,
		"einem",
	},
	{
		"kein",
		Der,
		true,
		CaseDative,
		"keinen",
	},
	{
		"unser",
		Das,
		false,
		CaseDative,
		"unserem",
	},
}

var definiteArticleCases = []struct {
	nounArticle Article
	isPlural    bool
	nounCase    Case
	result      string
}{
	{
		Das,
		false,
		CaseNominative,
		"das",
	},
	{
		Der,
		false,
		CaseAcusative,
		"den",
	},
	{
		Das,
		false,
		CaseDative,
		"dem",
	},
	{
		Die,
		false,
		CaseDative,
		"der",
	},
	{
		Der,
		true,
		CaseDative,
		"den",
	},
	{
		Das,
		false,
		CaseGenitive,
		"des",
	},
}
