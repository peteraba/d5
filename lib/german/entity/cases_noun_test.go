package entity

import "time"

var nounRegexpSuccessCases = []struct {
	raw, german, plural, genitive, isPlural string
}{
	{"Verabredung,~en", "Verabredung", "~en", "", ""},
	{"Anspruch,⍨e", "Anspruch", "⍨e", "", ""},
	{"Vereinigten Staaten von Amerika,- (pl)", "Vereinigten Staaten von Amerika", "- ", "", "(pl)"},
	{"Hintergeräusch,~e,~s/~es", "Hintergeräusch", "~e", "~s/~es", ""},
	{"CD-Brenner,~", "CD-Brenner", "~", "", ""},
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

var nounCreationSuccessCases = []struct {
	articles, german, english, third, user, learned, score, tags string
	noun                                                         *Noun
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
		&Noun{
			DefaultWord{
				"Hintergeräusch",
				[]Meaning{
					Meaning{"background noise", ""},
				},
				[]Meaning{
					Meaning{"háttérzaj", ""},
				},
				"noun",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"sound"},
				[]string{},
			},
			[]Article{Das},
			[]string{"~e"},
			[]string{"~s", "~es"},
			false,
		},
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
		&Noun{
			DefaultWord{
				"Jurastudium",
				[]Meaning{
					Meaning{"law studies", ""},
				},
				[]Meaning{
					Meaning{"jogi tanulmány", ""},
				},
				"noun",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"studies"},
				[]string{},
			},
			[]Article{Das},
			[]string{"Jurastudien"},
			[]string{},
			false,
		},
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
		&Noun{
			DefaultWord{
				"Nebel",
				[]Meaning{
					Meaning{"fog, mist, haze", ""},
					Meaning{"nebula", "astronomy"},
				},
				[]Meaning{
					Meaning{"köd", ""},
				},
				"noun",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"visible"},
				[]string{},
			},
			[]Article{Der},
			[]string{"~"},
			[]string{},
			false,
		},
	},
	{
		"e",
		"Vereinigten Staaten von Amerika,- (pl)",
		"United States of America",
		"Amerikai Egyesült Államok",
		"peteraba",
		"2015-05-03",
		"5",
		"country",
		&Noun{
			DefaultWord{
				"Vereinigten Staaten von Amerika",
				[]Meaning{
					Meaning{"United States of America", ""},
				},
				[]Meaning{
					Meaning{"Amerikai Egyesült Államok", ""},
				},
				"noun",
				"peteraba",
				time.Date(2015, 5, 3, 0, 0, 0, 0, time.UTC),
				5,
				[]string{"country"},
				[]string{},
			},
			[]Article{Die},
			[]string{"-"},
			[]string{},
			true,
		},
	},
}

var getPluralCases = []struct {
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
			[]string{"~s", "~e"},
			[]string{},
			false,
		},
		[]string{"Gulaschs", "Gulasche"},
		1,
		"Gulaschs",
	},
	{
		Noun{
			DefaultWord{
				"Klamotten",
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
			[]string{},
			true,
		},
		[]string{"Klamotten"},
		1,
		"Klamotten",
	},
	{
		Noun{
			DefaultWord{
				"Jurastudium",
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
			[]string{"Jurastudien"},
			[]string{},
			false,
		},
		[]string{"Jurastudien"},
		1,
		"Jurastudien",
	},
	{
		Noun{
			DefaultWord{
				"Knast",
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
			[]string{"⍨e"},
			[]string{},
			false,
		},
		[]string{"Knäste"},
		1,
		"Knäste",
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

var getNounDeclensionCases = []struct {
	noun               Noun
	singularNominative []string
	singularAcusative  []string
	singularDative     []string
	singularGenitive   []string
	pluralNominative   []string
	pluralAcusative    []string
	pluralDative       []string
	pluralGenitive     []string
}{
	{
		Noun{
			DefaultWord{
				"Berg",
				[]Meaning{},
				[]Meaning{},
				"",
				"",
				time.Now(),
				5,
				[]string{},
				[]string{},
			},
			[]Article{Der},
			[]string{"~e"},
			[]string{"~s", "~es"},
			false,
		},
		[]string{"Berg"},
		[]string{"Berg"},
		[]string{"Berg"},
		[]string{"Bergs", "Berges"},
		[]string{"Berge"},
		[]string{"Berge"},
		[]string{"Bergen"},
		[]string{"Berge"},
	},
}
