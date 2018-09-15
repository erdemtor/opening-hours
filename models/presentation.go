package models

import (
	"fmt"
	"strings"
)

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
