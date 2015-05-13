package entity

import "time"

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
	// Only one parantheses is allowed for a meaning
	"to get sth. down (stairs)   (to heaven) ",
	// No words are allowed after the parantheses
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

var verbRegexpSuccessCases = []struct {
	raw, german, arguments string
}{
	{"wütend sein, bin, bist + auf (A)", "wütend sein, bin, bist ", "+ auf (A)"},
	{"absprechen, absprach, abgesprochen, absprichst, abspricht + sich (A) + über (A)", "absprechen, absprach, abgesprochen, absprichst, abspricht ", "+ sich (A) + über (A)"},
	{"abhauen, abhiebe/abhaute, abgehauen/abgehaut", "abhauen, abhiebe/abhaute, abgehauen/abgehaut", ""},
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
	errors                                                       []string
	word                                                         *Any
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
		[]string{},
		&Any{
			DefaultWord{
				"Ich versteh nur Bahnhof",
				[]Meaning{
					Meaning{"I understand just train-station", ""},
				},
				[]Meaning{
					Meaning{"Én csak a vasútállomásokat értem", ""},
				},
				"exp",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"idiom", "ithinkispider.com"},
				[]string{},
			},
		},
	},
}

var adjectiveCreationSuccessCases = []struct {
	german, english, third, user, learned, score, tags string
	adjective                                          *Adjective
}{
	{
		"ägyptisch",
		"Egyptian",
		"egyiptomi",
		"peteraba",
		"2015-05-03",
		"5",
		"object_from",
		&Adjective{
			DefaultWord{
				"ägyptisch",
				[]Meaning{
					Meaning{"Egyptian", ""},
				},
				[]Meaning{
					Meaning{"egyiptomi", ""},
				},
				"adjective",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"object_from"},
				[]string{},
			},
			[]string{},
			[]string{},
		},
	},
	{
		"andauernd,-",
		"continuous; ongoing",
		"folyamatos",
		"peteraba",
		"2015-05-03",
		"5",
		"",
		&Adjective{
			DefaultWord{
				"andauernd",
				[]Meaning{
					Meaning{"continuous", ""},
					Meaning{"ongoing", ""},
				},
				[]Meaning{
					Meaning{"folyamatos", ""},
				},
				"adjective",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]string{"-"},
			[]string{},
		},
	},
	{
		"aufmerksam",
		"alert, observant (animal); kind, nice",
		"éber (állat); figyelmes, előzékeny",
		"peteraba",
		"2015-05-03",
		"5",
		"person, animal",
		&Adjective{
			DefaultWord{
				"aufmerksam",
				[]Meaning{
					Meaning{"alert, observant", "animal"},
					Meaning{"kind, nice", ""},
				},
				[]Meaning{
					Meaning{"éber", "állat"},
					Meaning{"figyelmes, előzékeny", ""},
				},
				"adjective",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"person", "animal"},
				[]string{},
			},
			[]string{},
			[]string{},
		},
	},
	{
		"jung,⍨er,⍨sten",
		"junior",
		"kezdő",
		"peteraba",
		"2015-05-03",
		"5",
		"person",
		&Adjective{
			DefaultWord{
				"jung",
				[]Meaning{
					Meaning{"junior", ""},
				},
				[]Meaning{
					Meaning{"kezdő", ""},
				},
				"adjective",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"person"},
				[]string{},
			},
			[]string{"⍨er"},
			[]string{"⍨sten"},
		},
	},
	{
		"schmal,~er/⍨er,~sten/⍨sten",
		"narrow",
		"keskeny, szűk",
		"peteraba",
		"2015-05-03",
		"5",
		"room, clothes",
		&Adjective{
			DefaultWord{
				"schmal",
				[]Meaning{
					Meaning{"narrow", ""},
				},
				[]Meaning{
					Meaning{"keskeny, szűk", ""},
				},
				"adjective",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"room", "clothes"},
				[]string{},
			},
			[]string{"~er", "⍨er"},
			[]string{"~sten", "⍨sten"},
		},
	},
}

var verbCreationSuccessCases = []struct {
	auxiliary, german, english, third, user, learned, score, tags string
	verb                                                          *Verb
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
		&Verb{
			DefaultWord{
				"brechen",
				[]Meaning{
					Meaning{"to break, to get broken", ""},
				},
				[]Meaning{
					Meaning{"összetörni, összetörik", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Haben},
			Prefix{},
			"",
			"",
			[]string{"brach"},
			[]string{"gebrochen"},
			[]string{},
			[]string{"brichst"},
			[]string{"bricht"},
			[]string{"brechen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
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
		&Verb{
			DefaultWord{
				"durchfallen",
				[]Meaning{
					Meaning{"to check through", ""},
				},
				[]Meaning{
					Meaning{"ellenőrizni, átvizsgálni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Haben},
			Prefix{
				"durch",
				true,
			},
			"",
			"",
			[]string{"durchfiel"},
			[]string{"durchgefallen"},
			[]string{},
			[]string{"durchfällst"},
			[]string{"durchfällt"},
			[]string{"durchfallen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
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
		&Verb{
			DefaultWord{
				"einfallen",
				[]Meaning{
					Meaning{"to come to mind", ""},
					Meaning{"to remember", ""},
				},
				[]Meaning{
					Meaning{"eszébe jut", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Sein},
			Prefix{
				"ein",
				true,
			},
			"",
			"",
			[]string{"einfiel"},
			[]string{"eingefallen"},
			[]string{},
			[]string{"einfällst"},
			[]string{"einfällt"},
			[]string{"einfallen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{
				Argument{"", CaseDative},
			},
		},
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
		&Verb{
			DefaultWord{
				"fehlen",
				[]Meaning{
					Meaning{"to lack", ""},
				},
				[]Meaning{
					Meaning{"hiányozni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Haben},
			Prefix{},
			"",
			"",
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{"fehlen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{
				Argument{"", CaseDative},
				Argument{"an", CaseDative},
			},
		},
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
		&Verb{
			DefaultWord{
				"tun",
				[]Meaning{
					Meaning{"to do", ""},
				},
				[]Meaning{
					Meaning{"tenni, csinálni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Haben},
			Prefix{},
			"",
			"",
			[]string{"taten"},
			[]string{"getan"},
			[]string{"tue"},
			[]string{"tust"},
			[]string{"tut"},
			[]string{"tun"},
			[]string{"tut"},
			[]string{"tun"},
			ReflexiveWithout,
			[]Argument{},
		},
	},
	{
		"h",
		"Sport treiben, trieb, getrieben",
		"to sport",
		"sportolni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
		&Verb{
			DefaultWord{
				"Sport treiben",
				[]Meaning{
					Meaning{"to sport", ""},
				},
				[]Meaning{
					Meaning{"sportolni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Haben},
			Prefix{},
			"Sport",
			"",
			[]string{"trieb"},
			[]string{"getrieben"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"treiben"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
	},
	{
		"s",
		"nett sein, bin, bist, ist, sein, seid, sein, war, gewesen + zu (D)",
		"to sneeze",
		"tüsszenteni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
		&Verb{
			DefaultWord{
				"nett sein",
				[]Meaning{
					Meaning{"to sneeze", ""},
				},
				[]Meaning{
					Meaning{"tüsszenteni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
			},
			[]Auxiliary{Sein},
			Prefix{},
			"",
			"nett",
			[]string{"war"},
			[]string{"gewesen"},
			[]string{"bin"},
			[]string{"bist"},
			[]string{"ist"},
			[]string{"sein"},
			[]string{"seid"},
			[]string{"sein"},
			ReflexiveWithout,
			[]Argument{
				Argument{"zu", "D"},
			},
		},
	},
}

var getGenitiveCases = []struct {
	noun            Noun
	expectedResult  []string
	stringCount     int
	expectedResult2 string
}{
	{
		Noun{
			DefaultWord{
				"Gulasch",
				[]Meaning{},
				[]Meaning{},
				"",
				"",
				time.Now(),
				5,
				[]string{},
				[]string{},
			},
			[]Article{},
			[]string{},
			[]string{"~es", "~s"},
			false,
		},
		[]string{"Gulasches", "Gulaschs"},
		1,
		"Gulasches",
	},
}
