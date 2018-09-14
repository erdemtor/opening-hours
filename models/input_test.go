package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpeningHours_ConvertToTimeline(t *testing.T) {
	var openingHoursInput = OpeningHours{
		"Monday":    Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Tuesday":   Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Wednesday": Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Thursday":  Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Friday":    Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Saturday":  Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
		"Sunday":    Events{Event{Type: Open, Value: 3600}, Event{Type: Close, Value: 36000}},
	}
	timeline, err := openingHoursInput.ConvertToTimeline()
	assert.NoError(t, err)
	for i := 0; i < len(timeline)-1; i++ {
		duration := timeline[i]
		assert.True(t, duration.Open.Before(duration.Close))
		assert.EqualValues(t, 36000-3600, duration.Close.Sub(duration.Open).Seconds())
	}
}
