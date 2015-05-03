package german

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

var auxiliaryRegexpSuccessCases = []struct {
	raw, first, second string
}{
	{"s", "s", ""},
	{"s/h", "s", "h"},
}

var auxiliaryRegexpFailureCases = []string{
	"",
	"i",
	"s/h/s",
	"S",
}

var argumentRegexpSuccessCases = []struct {
	raw, argPrep, argCase string
}{
	{" für (A)", " für ", "A"},
	{" bei (D)", " bei ", "D"},
}

var argumentRegexpFailureCases = []string{
	"",
	" bei (K)",
	" bei (D) bei (D)",
}

var meaningRegexpSuccessCases = []struct {
	raw, main, parant string
}{
	{"to get sth", "to get sth", ""},
	{"to get sth. down ", "to get sth. down ", ""},
	{"to get sth. down (stairs)", "to get sth. down ", "stairs"},
	{"to get sth. down (stairs)    ", "to get sth. down ", "stairs"},
}

var meaningRegexpFailureCases = []string{
	"",
	"to get sth. down (stairs)   (to heaven) ",
}

var nounRegexpSuccessCases = []struct {
	raw, german, plural string
}{
	{"Verabredung,~en", "Verabredung", "~en"},
	{"Anspruch,⍨e", "Anspruch", "⍨e"},
	{"Vereinigten Staaten von Amerika,- (pl)", "Vereinigten Staaten von Amerika", "- (pl)"},
}

var nounRegexpFailureCases = []string{
	"",
	"Verabrendung",
	"Verabrendung,!s/h/s",
	"Verabrendung,as,is,as",
	"kőr",
}

var adjectiveRegexpSuccessCases = []struct {
	raw, german, comparative, superlative string
}{
	{"erfolgsverwöhnt", "erfolgsverwöhnt", "", ""},
	{"erfahren,-", "erfahren", "-", ""},
	{"jung,⍨er,⍨sten", "jung", "⍨er", "⍨sten"},
}

var adjectiveRegexpFailureCases = []string{
	"",
	"Erfolgsverwohnt",
	"erfahren,-,-,-",
	"rót",
}

var verbRegexpSuccessCases = []struct {
	raw, german, arguments string
}{
	{"wütend sein, bin, bist + auf (A)", "wütend sein, bin, bist ", "+ auf (A)"},
	{"absprechen, absprach, abgesprochen, absprichst, abspricht + sich (A) + über (A)", "absprechen, absprach, abgesprochen, absprichst, abspricht ", "+ sich (A) + über (A)"},
}

var verbRegexpFailureCases = []string{
	"",
	"wűtend sein",
}

var argumentCreationCases = []struct {
	allArguments string
	arguments    []Argument
}{
	{
		"+ über (A)",
		[]Argument{
			Argument{"über", "A"},
		},
	},
	{
		"+ (D) + an (D)",
		[]Argument{
			Argument{"", "D"},
			Argument{"an", "D"},
		},
	},
}

var meaningCreationCases = []struct {
	allMeanings string
	meanings    []Meaning
}{
	{
		"to spank, to beat, to hit (colloquial)",
		[]Meaning{
			Meaning{"to spank, to beat, to hit", "colloquial"},
		},
	},
	{
		"beleolvasni (átv. is); sorok között olvasni",
		[]Meaning{
			Meaning{"beleolvasni", "átv. is"},
			Meaning{"sorok között olvasni", ""},
		},
	},
}

var wordCreationSuccessCases = []struct {
	german, english, third, category, user, learned, score, tags string
	ok                                                           bool
}{
	{
		"Ich versteh nur Bahnhof",
		"I understand just train-station",
		"Én csak a vasútállomásokat értem",
		"exp",
		"peteraba",
		"2015-05-03",
		"5",
		"idiom, ithinkispider.com",
		true,
	},
}
