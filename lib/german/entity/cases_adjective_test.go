package entity

import (
	"time"

	"github.com/peteraba/d5/lib/general"
)

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
				"",
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
				[]*general.Score{},
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
				"",
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
				[]*general.Score{},
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
				"",
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
				[]*general.Score{},
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
				"",
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
				[]*general.Score{},
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
				"",
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
				[]*general.Score{},
			},
			[]string{"~er", "⍨er"},
			[]string{"~sten", "⍨sten"},
		},
	},
}

var adjectiveCreationFailureCases = []struct {
	german, english, third, user, learned, score, tags string
}{
	{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	},
}

var getComparativeCases = []struct {
	adjective       Adjective
	expectedResult  []string
	maxCount        int
	expectedResult2 string
}{
	{
		Adjective{
			DefaultWord{
				"",
				"jung",
				[]Meaning{},
				[]Meaning{},
				"",
				"",
				time.Now(),
				5,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]string{"⍨er"},
			[]string{"⍨sten"},
		},
		[]string{"jünger"},
		1,
		"jünger",
	},
}

var getSuperlativeCases = []struct {
	adjective       Adjective
	expectedResult  []string
	maxCount        int
	expectedResult2 string
}{
	{
		Adjective{
			DefaultWord{
				"",
				"jung",
				[]Meaning{},
				[]Meaning{},
				"",
				"",
				time.Now(),
				5,
				[]string{},
				[]string{},
				[]*general.Score{},
			},
			[]string{"⍨er"},
			[]string{"⍨sten"},
		},
		[]string{"jüngsten"},
		1,
		"jüngsten",
	},
}

var declineAdjective = Adjective{
	DefaultWord{
		"",
		"neu",
		[]Meaning{},
		[]Meaning{},
		"",
		"",
		time.Now(),
		5,
		[]string{},
		[]string{},
		[]*general.Score{},
	},
	[]string{"~er"},
	[]string{"~sten"},
}

var declineAdjectiveCases = []struct {
	adjective      Adjective
	degree         Degree
	declension     Declension
	nounArticle    Article
	isPlural       bool
	nounCase       Case
	expectedResult string
}{
	{declineAdjective, Positive, Strong, Der, false, CaseNominative, "neuer"},
	{declineAdjective, Positive, Strong, Der, false, CaseAcusative, "neuen"},
	{declineAdjective, Positive, Strong, Der, false, CaseDative, "neuem"},
	{declineAdjective, Positive, Strong, Der, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Strong, Das, false, CaseNominative, "neues"},
	{declineAdjective, Positive, Strong, Das, false, CaseAcusative, "neues"},
	{declineAdjective, Positive, Strong, Das, false, CaseDative, "neuem"},
	{declineAdjective, Positive, Strong, Das, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Strong, Die, false, CaseNominative, "neue"},
	{declineAdjective, Positive, Strong, Die, false, CaseAcusative, "neue"},
	{declineAdjective, Positive, Strong, Die, false, CaseDative, "neuer"},
	{declineAdjective, Positive, Strong, Die, false, CaseGenitive, "neuer"},
	{declineAdjective, Positive, Strong, Der, true, CaseNominative, "neue"},
	{declineAdjective, Positive, Strong, Der, true, CaseAcusative, "neue"},
	{declineAdjective, Positive, Strong, Der, true, CaseDative, "neuen"},
	{declineAdjective, Positive, Strong, Der, true, CaseGenitive, "neuer"},

	{declineAdjective, Positive, Mixed, Der, false, CaseNominative, "neuer"},
	{declineAdjective, Positive, Mixed, Der, false, CaseAcusative, "neuen"},
	{declineAdjective, Positive, Mixed, Der, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Mixed, Der, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Mixed, Das, false, CaseNominative, "neues"},
	{declineAdjective, Positive, Mixed, Das, false, CaseAcusative, "neues"},
	{declineAdjective, Positive, Mixed, Das, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Mixed, Das, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Mixed, Die, false, CaseNominative, "neue"},
	{declineAdjective, Positive, Mixed, Die, false, CaseAcusative, "neue"},
	{declineAdjective, Positive, Mixed, Die, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Mixed, Die, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Mixed, Der, true, CaseNominative, "neuen"},
	{declineAdjective, Positive, Mixed, Der, true, CaseAcusative, "neuen"},
	{declineAdjective, Positive, Mixed, Der, true, CaseDative, "neuen"},
	{declineAdjective, Positive, Mixed, Der, true, CaseGenitive, "neuen"},

	{declineAdjective, Positive, Weak, Der, false, CaseNominative, "neue"},
	{declineAdjective, Positive, Weak, Der, false, CaseAcusative, "neuen"},
	{declineAdjective, Positive, Weak, Der, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Weak, Der, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Weak, Das, false, CaseNominative, "neue"},
	{declineAdjective, Positive, Weak, Das, false, CaseAcusative, "neue"},
	{declineAdjective, Positive, Weak, Das, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Weak, Das, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Weak, Die, false, CaseNominative, "neue"},
	{declineAdjective, Positive, Weak, Die, false, CaseAcusative, "neue"},
	{declineAdjective, Positive, Weak, Die, false, CaseDative, "neuen"},
	{declineAdjective, Positive, Weak, Die, false, CaseGenitive, "neuen"},
	{declineAdjective, Positive, Weak, Der, true, CaseNominative, "neuen"},
	{declineAdjective, Positive, Weak, Der, true, CaseAcusative, "neuen"},
	{declineAdjective, Positive, Weak, Der, true, CaseDative, "neuen"},
	{declineAdjective, Positive, Weak, Der, true, CaseGenitive, "neuen"},

	{declineAdjective, Comparative, Strong, Der, false, CaseNominative, "neuerer"},
	{declineAdjective, Comparative, Strong, Der, false, CaseAcusative, "neueren"},
	{declineAdjective, Comparative, Strong, Der, false, CaseDative, "neuerem"},
	{declineAdjective, Comparative, Strong, Der, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Strong, Das, false, CaseNominative, "neueres"},
	{declineAdjective, Comparative, Strong, Das, false, CaseAcusative, "neueres"},
	{declineAdjective, Comparative, Strong, Das, false, CaseDative, "neuerem"},
	{declineAdjective, Comparative, Strong, Das, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Strong, Die, false, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Strong, Die, false, CaseAcusative, "neuere"},
	{declineAdjective, Comparative, Strong, Die, false, CaseDative, "neuerer"},
	{declineAdjective, Comparative, Strong, Die, false, CaseGenitive, "neuerer"},
	{declineAdjective, Comparative, Strong, Der, true, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Strong, Der, true, CaseAcusative, "neuere"},
	{declineAdjective, Comparative, Strong, Der, true, CaseDative, "neueren"},
	{declineAdjective, Comparative, Strong, Der, true, CaseGenitive, "neuerer"},

	{declineAdjective, Comparative, Mixed, Der, false, CaseNominative, "neuerer"},
	{declineAdjective, Comparative, Mixed, Der, false, CaseAcusative, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Mixed, Das, false, CaseNominative, "neueres"},
	{declineAdjective, Comparative, Mixed, Das, false, CaseAcusative, "neueres"},
	{declineAdjective, Comparative, Mixed, Das, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Mixed, Das, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Mixed, Die, false, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Mixed, Die, false, CaseAcusative, "neuere"},
	{declineAdjective, Comparative, Mixed, Die, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Mixed, Die, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, true, CaseNominative, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, true, CaseAcusative, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, true, CaseDative, "neueren"},
	{declineAdjective, Comparative, Mixed, Der, true, CaseGenitive, "neueren"},

	{declineAdjective, Comparative, Weak, Der, false, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Weak, Der, false, CaseAcusative, "neueren"},
	{declineAdjective, Comparative, Weak, Der, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Weak, Der, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Weak, Das, false, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Weak, Das, false, CaseAcusative, "neuere"},
	{declineAdjective, Comparative, Weak, Das, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Weak, Das, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Weak, Die, false, CaseNominative, "neuere"},
	{declineAdjective, Comparative, Weak, Die, false, CaseAcusative, "neuere"},
	{declineAdjective, Comparative, Weak, Die, false, CaseDative, "neueren"},
	{declineAdjective, Comparative, Weak, Die, false, CaseGenitive, "neueren"},
	{declineAdjective, Comparative, Weak, Der, true, CaseNominative, "neueren"},
	{declineAdjective, Comparative, Weak, Der, true, CaseAcusative, "neueren"},
	{declineAdjective, Comparative, Weak, Der, true, CaseDative, "neueren"},
	{declineAdjective, Comparative, Weak, Der, true, CaseGenitive, "neueren"},

	{declineAdjective, Superlative, Strong, Der, false, CaseNominative, "neustener"},
	{declineAdjective, Superlative, Strong, Der, false, CaseAcusative, "neustenen"},
	{declineAdjective, Superlative, Strong, Der, false, CaseDative, "neustenem"},
	{declineAdjective, Superlative, Strong, Der, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Strong, Das, false, CaseNominative, "neustenes"},
	{declineAdjective, Superlative, Strong, Das, false, CaseAcusative, "neustenes"},
	{declineAdjective, Superlative, Strong, Das, false, CaseDative, "neustenem"},
	{declineAdjective, Superlative, Strong, Das, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Strong, Die, false, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Strong, Die, false, CaseAcusative, "neustene"},
	{declineAdjective, Superlative, Strong, Die, false, CaseDative, "neustener"},
	{declineAdjective, Superlative, Strong, Die, false, CaseGenitive, "neustener"},
	{declineAdjective, Superlative, Strong, Der, true, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Strong, Der, true, CaseAcusative, "neustene"},
	{declineAdjective, Superlative, Strong, Der, true, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Strong, Der, true, CaseGenitive, "neustener"},

	{declineAdjective, Superlative, Mixed, Der, false, CaseNominative, "neustener"},
	{declineAdjective, Superlative, Mixed, Der, false, CaseAcusative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Mixed, Das, false, CaseNominative, "neustenes"},
	{declineAdjective, Superlative, Mixed, Das, false, CaseAcusative, "neustenes"},
	{declineAdjective, Superlative, Mixed, Das, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Das, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Mixed, Die, false, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Mixed, Die, false, CaseAcusative, "neustene"},
	{declineAdjective, Superlative, Mixed, Die, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Die, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, true, CaseNominative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, true, CaseAcusative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, true, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Mixed, Der, true, CaseGenitive, "neustenen"},

	{declineAdjective, Superlative, Weak, Der, false, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Weak, Der, false, CaseAcusative, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Weak, Das, false, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Weak, Das, false, CaseAcusative, "neustene"},
	{declineAdjective, Superlative, Weak, Das, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Weak, Das, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Weak, Die, false, CaseNominative, "neustene"},
	{declineAdjective, Superlative, Weak, Die, false, CaseAcusative, "neustene"},
	{declineAdjective, Superlative, Weak, Die, false, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Weak, Die, false, CaseGenitive, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, true, CaseNominative, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, true, CaseAcusative, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, true, CaseDative, "neustenen"},
	{declineAdjective, Superlative, Weak, Der, true, CaseGenitive, "neustenen"},
}
