package models

import (
	"errors"
	"fmt"
	"time"
)

type OpeningHours map[string]Events

type Events []Event

type Event struct {
	Type  eventType `json:"type"`
	Value int64     `json:"value"`
	Time  time.Time `json:"-"`
}

type eventType string

const (
	Open  eventType = "open"
	Close eventType = "close"
)

func (o OpeningHours) ConvertToTimeline() (Timeline, error) {
	events, err := o.toEventStreamWithCorrectedDateValues()
	if err != nil {
		return Timeline{}, nil
	}
	var timeline Timeline
	for i := 0; i < len(events); i += 2 {
		eventOpen := events[i]
		eventClose := events[i+1]
		if eventOpen.Type != Open || eventClose.Type != Close {
			return timeline, errors.New("open and close events are not given in order")
		}
		timeline = append(timeline, OpenDuration{Open: eventOpen.Time, Close: eventClose.Time})
	}
	return timeline, nil
}

func (o OpeningHours) toEventStreamWithCorrectedDateValues() (Events, error) {
	var allEvents Events
	for dayIndex, day := range weekdaysSorted {
		events, ok := o[day]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Day: %s is not found in the opening hours", day))
		}
		for _, event := range events {
			event.Time = time.Unix(event.Value, 0).AddDate(0, 0, dayIndex-3).UTC()
			allEvents = append(allEvents, event)
		}
	}
	return allEvents, nil
}
