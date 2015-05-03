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
	// Only one paranthases is allowed for a meaning
	"to get sth. down (stairs)   (to heaven) ",
	// No words are allowed after the paranthases
	"to get sth. down (stairs)   to heaven ",
}

var adjectiveRegexpSuccessCases = []struct {
	raw, german, comparative, superlative string
}{
	{"erfolgsverwöhnt", "erfolgsverwöhnt", "", ""},
	{"erfahren,-", "erfahren", "-", ""},
	{"jung,⍨er,⍨sten", "jung", "⍨er", "⍨sten"},
}

var adjectiveRegexpFailureCases = []string{
	// Empty string is not a valid adjective
	"",
	// Adjectives must be all lower-cased
	"Erfolgsverwohnt",
	// Adjective can only contain one commas
	"erfahren,-,-,-",
	// Only German alphabet is allowed
	"rót",
}

var nounRegexpSuccessCases = []struct {
	raw, german, plural, genitive, isPlural string
}{
	{"Verabredung,~en", "Verabredung", "~en", "", ""},
	{"Anspruch,⍨e", "Anspruch", "⍨e", "", ""},
	{"Vereinigten Staaten von Amerika,- (pl)", "Vereinigten Staaten von Amerika", "- ", "", "(pl)"},
	{"Hintergeräusch,~e,~s/~es", "Hintergeräusch", "~e", "~s/~es", ""},
}

var nounRegexpFailureCases = []string{
	// Empty string is not a valid noun
	"",
	// A noun should not have more than two commas
	"Hintergeräusch,~e,~s/~e,~s/~es",
	// Exclaimation is not a valid word-character
	"Hintergeräusch,!~e,~s/~es",
	// Only German alphabet allowed
	"kőr",
}

var verbRegexpSuccessCases = []struct {
	raw, german, arguments string
}{
	{"wütend sein, bin, bist + auf (A)", "wütend sein, bin, bist ", "+ auf (A)"},
	{"absprechen, absprach, abgesprochen, absprichst, abspricht + sich (A) + über (A)", "absprechen, absprach, abgesprochen, absprichst, abspricht ", "+ sich (A) + über (A)"},
}

var verbRegexpFailureCases = []string{
	// Empty string is not a valid verb
	"",
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

var adjectiveCreationSuccessCases = []struct {
	german, english, third, user, learned, score, tags string
}{
	{
		"ägyptisch",
		"Egyptian",
		"egyiptomi",
		"peteraba",
		"2015-05-03",
		"5",
		"object_from",
	},
	{
		"andauernd,-",
		"continuous; ongoing",
		"folyamatos",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"aufmerksam",
		"alert, observant (animal); kind, nice",
		"éber (állat); figyelmes, előzékeny",
		"peteraba",
		"2015-05-03",
		"5",
		"person, animal",
	},
	{
		"jung,⍨er,⍨sten",
		"junior",
		"kezdő",
		"peteraba",
		"2015-05-03",
		"5",
		"person",
	},
	{
		"schmal,~er/⍨er,~sten/⍨sten",
		"narrow",
		"keskeny, szűk",
		"peteraba",
		"2015-05-03",
		"5",
		"room, clothes",
	},
}

var nounCreationSuccessCases = []struct {
	articles, german, english, third, user, learned, score, tags string
}{
	{
		"s",
		"Hintergeräusch,~e,~s/~es",
		"background noise",
		"háttérzaj",
		"peteraba",
		"2015-05-03",
		"5",
		"sound",
	},
	{
		"s",
		"Jurastudium, Jurastudien",
		"law studies",
		"jogi tanulmány",
		"peteraba",
		"2015-05-03",
		"5",
		"studies",
	},
	{
		"r",
		"Nebel,~",
		"fog, mist, haze; nebula (astronomy)",
		"köd",
		"peteraba",
		"2015-05-03",
		"5",
		"visible",
	},
}

var verbCreationSuccessCases = []struct {
	auxiliary, german, english, third, user, learned, score, tags string
}{
	{
		"h",
		"brechen, brach, gebrochen, brichst, bricht",
		"to break, to get broken",
		"összetörni, összetörik",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"durch|fallen, durchfiel, durchgefallen, durchfällst, durchfällt",
		"to check through",
		"ellenőrizni, átvizsgálni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"s",
		"einfallen, einfiel, eingefallen, einfällst, einfällt + (D)",
		"to come to mind; to remember",
		"eszébe jut",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"fehlen + (D) + an (D)",
		"to lack",
		"hiányozni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"tun, tue, tust, tut, tun, tut, tun, taten, getan",
		"to do",
		"tenni, csinálni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
}
