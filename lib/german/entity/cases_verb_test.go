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

var verbCreationSuccessCases = []struct {
	auxiliary, german, english, third, user, learned, score, tags    string
	verb                                                             *Verb
	presentS1, presentS2, presentS3, presentP1, presentP2, presentP3 []string
	pastS1, pastS2, pastS3, pastP1, pastP2, pastP3                   []string
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
			[]string{"gebrochen"},
			[]string{"brach"},
			[]string{},
			[]string{"brichst"},
			[]string{"bricht"},
			[]string{"brechen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
		[]string{"breche"},
		[]string{"brichst"},
		[]string{"bricht"},
		[]string{"brechen"},
		[]string{"brecht"},
		[]string{"brechen"},
		[]string{"brach"},
		[]string{"brachst"},
		[]string{"brach"},
		[]string{"brachen"},
		[]string{"bracht"},
		[]string{"brachen"},
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
			[]string{"durchgefallen"},
			[]string{"durchfiel"},
			[]string{},
			[]string{"durchfällst"},
			[]string{"durchfällt"},
			[]string{"durchfallen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
		[]string{"durchfalle"},
		[]string{"durchfällst"},
		[]string{"durchfällt"},
		[]string{"durchfallen"},
		[]string{"durchfallt"},
		[]string{"durchfallen"},
		[]string{"durchfiel"},
		[]string{"durchfielst"},
		[]string{"durchfiel"},
		[]string{"durchfielen"},
		[]string{"durchfielt"},
		[]string{"durchfielen"},
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
			[]string{"eingefallen"},
			[]string{"einfiel"},
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
		[]string{"einfalle"},
		[]string{"einfällst"},
		[]string{"einfällt"},
		[]string{"einfallen"},
		[]string{"einfallt"},
		[]string{"einfallen"},
		[]string{"einfiel"},
		[]string{"einfielst"},
		[]string{"einfiel"},
		[]string{"einfielen"},
		[]string{"einfielt"},
		[]string{"einfielen"},
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
		[]string{"fehle"},
		[]string{"fehlst"},
		[]string{"fehlt"},
		[]string{"fehlen"},
		[]string{"fehlt"},
		[]string{"fehlen"},
		[]string{"fehlte"},
		[]string{"fehltest"},
		[]string{"fehlte"},
		[]string{"fehlten"},
		[]string{"fehltet"},
		[]string{"fehlten"},
	},
	{
		"h",
		"tun, tue, tust, tut, tun, tut, tun, tat, getan",
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
			[]string{"getan"},
			[]string{"tat"},
			[]string{"tue"},
			[]string{"tust"},
			[]string{"tut"},
			[]string{"tun"},
			[]string{"tut"},
			[]string{"tun"},
			ReflexiveWithout,
			[]Argument{},
		},
		[]string{"tue"},
		[]string{"tust"},
		[]string{"tut"},
		[]string{"tun"},
		[]string{"tut"},
		[]string{"tun"},
		[]string{"tat"},
		[]string{"tatst"},
		[]string{"tat"},
		[]string{"taten"},
		[]string{"tatet"},
		[]string{"taten"},
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
			[]string{"getrieben"},
			[]string{"trieb"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"treiben"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
		},
		[]string{"treibe"},
		[]string{"treibst"},
		[]string{"treibt"},
		[]string{"treiben"},
		[]string{"treibt"},
		[]string{"treiben"},
		[]string{"trieb"},
		[]string{"triebt"},
		[]string{"trieb"},
		[]string{"trieben"},
		[]string{"triebt"},
		[]string{"trieben"},
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
			[]string{"gewesen"},
			[]string{"war"},
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
		[]string{"bin"},
		[]string{"bist"},
		[]string{"ist"},
		[]string{"sein"},
		[]string{"seid"},
		[]string{"sein"},
		[]string{"war"},
		[]string{"warst"},
		[]string{"war"},
		[]string{"waren"},
		[]string{"wart"},
		[]string{"waren"},
	},
}
