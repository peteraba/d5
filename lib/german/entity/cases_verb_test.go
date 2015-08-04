package entity

import (
	"time"

	"github.com/peteraba/d5/lib/general"
)

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

var prefixCreationCases = []struct {
	german string
	prefix Prefix
}{
	{
		"ausfüllen",
		Prefix{"aus", true},
	},
	{
		"verfüllen",
		Prefix{"ver", false},
	},
	{
		"ausfüllen",
		Prefix{"aus", true},
	},
	{
		"fehlen",
		Prefix{"", false},
	},
	{
		"durchsuchen",
		Prefix{"durch", false},
	},
	{
		"durch|rechnen",
		Prefix{"durch", true},
	},
	{
		"arbeiten",
		Prefix{"", false},
	},
}

var verbCreationSuccessCases = []struct {
	auxiliary, german, english, third, user, learned, score, tags                string
	verb                                                                         *Verb
	presentS1, presentS2, presentS3, presentP1, presentP2, presentP3             []string
	preteriteS1, preteriteS2, preteriteS3, preteriteP1, preteriteP2, preteriteP3 []string
	tense                                                                        Tense
	pp                                                                           PersonalPronoun
	expectedSeparated                                                            [][2]string
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
				[]*general.Score{},
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
			"",
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
		Present,
		S1,
		[][2]string{
			[2]string{"breche", ""},
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
				[]*general.Score{},
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
			"",
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
		Present,
		S2,
		[][2]string{
			[2]string{"fällst", "durch"},
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
				[]*general.Score{},
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
			"",
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
		Present,
		S3,
		[][2]string{
			[2]string{"fällt", "ein"},
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
				[]*general.Score{},
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
			"",
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
		Present,
		P1,
		[][2]string{
			[2]string{"fehlen", ""},
		},
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
				[]*general.Score{},
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
			"",
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
		Present,
		P2,
		[][2]string{
			[2]string{"tut", ""},
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
				"treiben",
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
				[]*general.Score{},
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
			"",
		},
		[]string{"treibe"},
		[]string{"treibst"},
		[]string{"treibt"},
		[]string{"treiben"},
		[]string{"treibt"},
		[]string{"treiben"},
		[]string{"trieb"},
		[]string{"triebst"},
		[]string{"trieb"},
		[]string{"trieben"},
		[]string{"triebt"},
		[]string{"trieben"},
		Present,
		P3,
		[][2]string{
			[2]string{"treiben", ""},
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
				"sein",
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
				[]*general.Score{},
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
			"",
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
		Preterite,
		S1,
		[][2]string{
			[2]string{"war", ""},
		},
	},
	{
		"s",
		"verzweifeln",
		"to panic",
		"kétségbeesni",
		"peteraba",
		"2015-05-03",
		"1",
		"",
		&Verb{
			DefaultWord{
				"verzweifeln",
				[]Meaning{
					Meaning{"to panic", ""},
				},
				[]Meaning{
					Meaning{"kétségbeesni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				1,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Sein},
			Prefix{"ver", false},
			"",
			"",
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{"verzweifeln"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
			"",
		},
		[]string{"verzweifele"},
		[]string{"verzweifelst"},
		[]string{"verzweifelt"},
		[]string{"verzweifeln"},
		[]string{"verzweifelt"},
		[]string{"verzweifeln"},
		[]string{"verzweifelte"},
		[]string{"verzweifeltest"},
		[]string{"verzweifelte"},
		[]string{"verzweifelten"},
		[]string{"verzweifeltet"},
		[]string{"verzweifelten"},
		Preterite,
		S2,
		[][2]string{
			[2]string{"verzweifeltest", ""},
		},
	},
	{
		"h",
		"bewegen, bewog/bewegte, bewogen/bewegt",
		"to persuade sb, to induce sb",
		"rábírni vkit (vmire)",
		"peteraba",
		"2015-05-08",
		"1",
		"",
		&Verb{
			DefaultWord{
				"bewegen",
				[]Meaning{
					Meaning{"to persuade sb, to induce sb", ""},
				},
				[]Meaning{
					Meaning{"rábírni vkit", "vmire"},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				1,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Haben},
			Prefix{"be", false},
			"",
			"",
			[]string{"bewogen", "bewegt"},
			[]string{"bewog", "bewegte"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"bewegen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{},
			"",
		},
		[]string{"bewege"},
		[]string{"bewegst"},
		[]string{"bewegt"},
		[]string{"bewegen"},
		[]string{"bewegt"},
		[]string{"bewegen"},
		[]string{"bewog", "bewegte"},
		[]string{"bewogst", "bewegtest"},
		[]string{"bewog", "bewegte"},
		[]string{"bewogen", "bewegten"},
		[]string{"bewogt", "bewegtet"},
		[]string{"bewogen", "bewegten"},
		Preterite,
		S3,
		[][2]string{
			[2]string{"bewog", ""},
			[2]string{"bewegte", ""},
		},
	},
	{
		"s",
		"einverstanden sein, bin, bist, ist, sein, seid, sein, war, gewesen",
		"to agree",
		"beleegyezni, egyetérteni",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"sein",
				[]Meaning{
					Meaning{"to agree", ""},
				},
				[]Meaning{
					Meaning{"beleegyezni, egyetérteni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Sein},
			Prefix{"", false},
			"",
			"einverstanden",
			[]string{"gewesen"},
			[]string{"war"},
			[]string{"bin"},
			[]string{"bist"},
			[]string{"ist"},
			[]string{"sein"},
			[]string{"seid"},
			[]string{"sein"},
			ReflexiveWithout,
			[]Argument{},
			"",
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
		Preterite,
		P1,
		[][2]string{
			[2]string{"waren", ""},
		},
	},
	{
		"s",
		"geschehen, -, -, geschieht, -, -, geschehen, geschah, geschehen",
		"to occur, to happen (formal)",
		"történni (formális)",
		"peteraba",
		"2015-05-08",
		"0",
		"",
		&Verb{
			DefaultWord{
				"geschehen",
				[]Meaning{
					Meaning{"to occur, to happen", "formal"},
				},
				[]Meaning{
					Meaning{"történni", "formális"},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				5,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Sein},
			Prefix{"ge", false},
			"",
			"",
			[]string{"geschehen"},
			[]string{"geschah"},
			[]string{"-"},
			[]string{"-"},
			[]string{"geschieht"},
			[]string{"-"},
			[]string{"-"},
			[]string{"geschehen"},
			ReflexiveWithout,
			[]Argument{},
			"",
		},
		[]string{"-"},
		[]string{"-"},
		[]string{"geschieht"},
		[]string{"-"},
		[]string{"-"},
		[]string{"geschehen"},
		[]string{"-"},
		[]string{"-"},
		[]string{"geschah"},
		[]string{"-"},
		[]string{"-"},
		[]string{"geschahen"},
		Preterite,
		P2,
		[][2]string{
			[2]string{"-", ""},
		},
	},
	{
		"s",
		"gleicher Meinung sein, bin, bist, ist, sein, seid, sein, war, gewesen",
		"to agree in sth, to have the same opinion on sth",
		"egyetérteni vmiben, egy véleményen lenni",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"sein",
				[]Meaning{
					Meaning{"to agree in sth, to have the same opinion on sth", ""},
				},
				[]Meaning{
					Meaning{"egyetérteni vmiben, egy véleményen lenni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Sein},
			Prefix{"", false},
			"Meinung",
			"gleicher",
			[]string{"gewesen"},
			[]string{"war"},
			[]string{"bin"},
			[]string{"bist"},
			[]string{"ist"},
			[]string{"sein"},
			[]string{"seid"},
			[]string{"sein"},
			ReflexiveWithout,
			[]Argument{},
			"",
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
		Preterite,
		P3,
		[][2]string{
			[2]string{"waren", ""},
		},
	},
	{
		"s",
		"sein, bin, -, -, sein, -, -, war, gewesen",
		"to be",
		"létezni",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"sein",
				[]Meaning{
					Meaning{"to be", ""},
				},
				[]Meaning{
					Meaning{"létezni", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Sein},
			Prefix{"", false},
			"",
			"",
			[]string{"gewesen"},
			[]string{"war"},
			[]string{"bin"},
			[]string{"-"},
			[]string{"-"},
			[]string{"sein"},
			[]string{"-"},
			[]string{"-"},
			ReflexiveWithout,
			[]Argument{},
			"",
		},
		[]string{"bin"},
		[]string{"-"},
		[]string{"-"},
		[]string{"sein"},
		[]string{"-"},
		[]string{"-"},
		[]string{"war"},
		[]string{"-"},
		[]string{"-"},
		[]string{"waren"},
		[]string{"-"},
		[]string{"-"},
		PastParticiple,
		P1,
		[][2]string{
			[2]string{"gewesen", ""},
		},
	},
	{
		"h",
		"ausgeben, ausgab, ausgegeben, ausgibst, ausgibt + sich (A) + als (N)",
		"to pose as sb, to personate sb",
		"kiadni magát vkinek",
		"peteraba",
		"2015-05-08",
		"2",
		"",
		&Verb{
			DefaultWord{
				"ausgeben",
				[]Meaning{
					Meaning{"to pose as sb, to personate sb", ""},
				},
				[]Meaning{
					Meaning{"kiadni magát vkinek", ""},
				},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				2,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Haben},
			Prefix{"aus", true},
			"",
			"",
			[]string{"ausgegeben"},
			[]string{"ausgab"},
			[]string{},
			[]string{"ausgibst"},
			[]string{"ausgibt"},
			[]string{"ausgeben"},
			[]string{},
			[]string{},
			ReflexiveAcusative,
			[]Argument{
				Argument{"als", CaseNominative},
			},
			"",
		},
		[]string{"ausgebe"},
		[]string{"ausgibst"},
		[]string{"ausgibt"},
		[]string{"ausgeben"},
		[]string{"ausgebt"},
		[]string{"ausgeben"},
		[]string{"ausgab"},
		[]string{"ausgabst"},
		[]string{"ausgab"},
		[]string{"ausgaben"},
		[]string{"ausgabt"},
		[]string{"ausgaben"},
		Preterite,
		P1,
		[][2]string{
			[2]string{"gaben", "aus"},
		},
	},
	{
		"h",
		"besinnen, besann, besonnen + sich (A) + (G)",
		"to think better of sth",
		"",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"besinnen",
				[]Meaning{
					Meaning{"to think better of sth", ""},
				},
				[]Meaning{},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Haben},
			Prefix{"be", false},
			"",
			"",
			[]string{"besonnen"},
			[]string{"besann"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"besinnen"},
			[]string{},
			[]string{},
			ReflexiveAcusative,
			[]Argument{
				Argument{"", CaseGenitive},
			},
			"",
		},
		[]string{"besinne"},
		[]string{"besinnst"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besann"},
		[]string{"besannst"},
		[]string{"besann"},
		[]string{"besannen"},
		[]string{"besannt"},
		[]string{"besannen"},
		PastParticiple,
		P1,
		[][2]string{
			[2]string{"besonnen", ""},
		},
	},
	{
		"h",
		"besinnen, besann, besonnen + sich (D) + (G)",
		"to think better of sth",
		"",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"besinnen",
				[]Meaning{
					Meaning{"to think better of sth", ""},
				},
				[]Meaning{},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]Auxiliary{Haben},
			Prefix{"be", false},
			"",
			"",
			[]string{"besonnen"},
			[]string{"besann"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"besinnen"},
			[]string{},
			[]string{},
			ReflexiveDative,
			[]Argument{
				Argument{"", CaseGenitive},
			},
			"",
		},
		[]string{"besinne"},
		[]string{"besinnst"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besann"},
		[]string{"besannst"},
		[]string{"besann"},
		[]string{"besannen"},
		[]string{"besannt"},
		[]string{"besannen"},
		PastParticiple,
		P1,
		[][2]string{
			[2]string{"besonnen", ""},
		},
	},
	{
		"h",
		"besinnen, besann, besonnen + sich (N) + (G)",
		"to think better of sth",
		"",
		"peteraba",
		"2015-05-08",
		"10",
		"",
		&Verb{
			DefaultWord{
				"besinnen",
				[]Meaning{
					Meaning{"to think better of sth", ""},
				},
				[]Meaning{},
				"verb",
				"peteraba",
				time.Date(2015, 5, 8, 0, 0, 0, 0, time.UTC),
				10,
				[]string{},
				[]string{"Reflexive definition is invalid"},
				[]*general.Score{},
			},
			[]Auxiliary{Haben},
			Prefix{"be", false},
			"",
			"",
			[]string{"besonnen"},
			[]string{"besann"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"besinnen"},
			[]string{},
			[]string{},
			ReflexiveWithout,
			[]Argument{
				Argument{"", CaseGenitive},
			},
			"",
		},
		[]string{"besinne"},
		[]string{"besinnst"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besinnt"},
		[]string{"besinnen"},
		[]string{"besann"},
		[]string{"besannst"},
		[]string{"besann"},
		[]string{"besannen"},
		[]string{"besannt"},
		[]string{"besannen"},
		PastParticiple,
		P1,
		[][2]string{
			[2]string{"besonnen", ""},
		},
	},
}

var verbCreationFailureCases = []struct {
	auxiliary, german, english, third, user, learned, score, tags string
}{
	{
		"h",
		"arbeiten, arbiet",
		"to work",
		"dolgozni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"arbeiten, arbiető",
		"to work",
		"dolgozni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"arbeiten, arbiet + an (B)",
		"to work",
		"dolgozni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
	{
		"h",
		"arbeiten, arbiet + aus / - hello",
		"to work",
		"dolgozni",
		"peteraba",
		"2015-05-03",
		"5",
		"",
	},
}
