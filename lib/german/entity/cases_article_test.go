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
	// Nominative
	{"ein", Der, false, CaseNominative, "ein"},
	{"ein", Die, false, CaseNominative, "eine"},
	{"ein", Das, false, CaseNominative, "ein"},
	{"ein", Der, true, CaseNominative, ""},
	{"kein", Der, true, CaseNominative, "keine"},
	// Acusative
	{"ein", Der, false, CaseAcusative, "einen"},
	{"ein", Die, false, CaseAcusative, "eine"},
	{"ein", Das, false, CaseAcusative, "ein"},
	{"ein", Der, true, CaseAcusative, ""},
	{"kein", Der, true, CaseAcusative, "keine"},
	// Dative
	{"ein", Der, false, CaseDative, "einem"},
	{"ein", Die, false, CaseDative, "einer"},
	{"ein", Das, false, CaseDative, "einem"},
	{"ein", Der, true, CaseDative, ""},
	{"kein", Der, true, CaseDative, "keinen"},
	// Genitive
	{"ein", Der, false, CaseGenitive, "eines"},
	{"ein", Die, false, CaseGenitive, "einer"},
	{"ein", Das, false, CaseGenitive, "eines"},
	{"ein", Der, true, CaseGenitive, ""},
	{"kein", Der, true, CaseGenitive, "keiner"},
	// Extra
	{"unser", Das, false, CaseDative, "unserem"},
	{"", Das, false, CaseDative, "einem"},
}

var definiteArticleCases = []struct {
	word        string
	nounArticle Article
	isPlural    bool
	nounCase    Case
	result      string
}{
	{"das", Der, false, CaseNominative, "der"},
	{"das", Die, false, CaseNominative, "die"},
	{"das", Das, false, CaseNominative, "das"},
	{"das", Das, true, CaseNominative, "die"},
	{"das", Der, false, CaseAcusative, "den"},
	{"das", Die, false, CaseAcusative, "die"},
	{"das", Das, false, CaseAcusative, "das"},
	{"das", Der, true, CaseAcusative, "die"},
	{"das", Der, false, CaseDative, "dem"},
	{"das", Die, false, CaseDative, "der"},
	{"das", Das, false, CaseDative, "dem"},
	{"das", Der, true, CaseDative, "den"},
	{"das", Der, false, CaseGenitive, "des"},
	{"das", Die, false, CaseGenitive, "der"},
	{"das", Das, false, CaseGenitive, "des"},
	{"das", Das, true, CaseGenitive, "der"},

	{"diese", Der, false, CaseNominative, "dieser"},
	{"diese", Die, false, CaseNominative, "diese"},
	{"diese", Das, false, CaseNominative, "dieses"},
	{"diese", Der, true, CaseNominative, "diese"},
	{"diese", Der, false, CaseAcusative, "diesen"},
	{"diese", Die, false, CaseAcusative, "diese"},
	{"diese", Das, false, CaseAcusative, "dieses"},
	{"diese", Der, true, CaseAcusative, "diese"},
	{"diese", Der, false, CaseDative, "diesem"},
	{"diese", Die, false, CaseDative, "dieser"},
	{"diese", Das, false, CaseDative, "diesem"},
	{"diese", Der, true, CaseDative, "diesen"},
	{"diese", Der, false, CaseGenitive, "dieses"},
	{"diese", Die, false, CaseGenitive, "dieser"},
	{"diese", Das, false, CaseGenitive, "dieses"},
	{"diese", Der, true, CaseGenitive, "dieser"},
}
