package models

import (
	"strings"
	"time"
)

//Timeline is array of open-durations representing the start-end times in which the restaurant is open
type Timeline []OpenDuration

//OpenDuration represents a time-slot in which the restaurant is open
type OpenDuration struct {
	Open  time.Time
	Close time.Time
}

//ToOutputFormat
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
