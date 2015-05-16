package util

import "time"

func ParseTimeNow(timeForm, toParse string) time.Time {
	timeParsed, err := time.Parse(timeForm, toParse)
	if err != nil {
		timeParsed = time.Now()
	}

	return timeParsed
}
