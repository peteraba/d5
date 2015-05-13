package entity

import "time"

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
