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
/*	{"erfolgsverwöhnt", "erfolgsverwöhnt", "", ""},
	{"erfahren,-", "erfahren", "-", ""},
	{"jung,⍨er,⍨sten", "jung", "⍨er", "⍨sten"},
*/
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
/*	{"wütend sein, bin, bist + auf (A)", "wütend sein, bin, bist ", "+ auf (A)"},
	{"absprechen, absprach, abgesprochen, absprichst, abspricht + sich (A) + über (A)", "absprechen, absprach, abgesprochen, absprichst, abspricht ", "+ sich (A) + über (A)"},
*/
}

var verbRegexpFailureCases = []string{
	"",
	"wűtend sein",
}
