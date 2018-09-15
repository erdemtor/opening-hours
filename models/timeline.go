package models

import (
	"fmt"
	"strings"
	"time"
)

type Timeline []OpenDuration

type OpenDuration struct {
	Open  time.Time
	Close time.Time
}
type stringTuple struct {
	first  string
	second string
}

func (s stringTuple) String() string {
	return strings.Join([]string{s.first, s.second}, " - ")
}

type stringTuples []stringTuple

func (s stringTuples) String() string {
	var stringTuplesStringfied []string
	for _, sTuple := range s {
		stringTuplesStringfied = append(stringTuplesStringfied, sTuple.String())
	}
	return strings.Join(stringTuplesStringfied, ", ")
}

type Presentation map[string]stringTuples

func (p Presentation) String() string {
	var lines []string
	for _, weekday := range weekdaysSorted {
		var line string
		tuples, exists := p[weekday]
		if !exists {
			line = fmt.Sprintf("%s: Closed", strings.Title(weekday))
		} else {
			line = fmt.Sprintf("%s: %s", strings.Title(weekday), tuples.String())
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
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
