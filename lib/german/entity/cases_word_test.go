package entity

import (
	"time"

	"github.com/peteraba/d5/lib/general"
)

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

var meaningCreationSuccessCases = []struct {
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

var meaningCreationFailureCases = []struct {
	allMeanings string
}{
	{
		"to spank, to beat, to hit (colloquial) (colloquial)",
	},
	{
		"beleolvasni (átv. is); sorok között olvasni <>!@#$^&*()~",
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
		"15",
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
				[]*general.Score{},
			},
			"",
		},
	},
}
