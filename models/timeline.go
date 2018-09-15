package models

import (
	"strings"
	"time"
)

type Timeline []OpenDuration

type OpenDuration struct {
	Open  time.Time
	Close time.Time
}

func (t Timeline) ToOutputFormat() Presentation {
	out := Presentation{}
	for _, openDuration := range t {
		day := strings.ToLower(openDuration.Open.Weekday().String())
		theSameDayIntervals, ok := out[day]
		if !ok {
			theSameDayIntervals = []stringTuple{}
		}
		theSameDayIntervals = append(theSameDayIntervals,
			stringTuple{first: openDuration.Open.Format(time.Kitchen),
				second: openDuration.Close.Format(time.Kitchen)})
		out[day] = theSameDayIntervals
	}
	return out
}
