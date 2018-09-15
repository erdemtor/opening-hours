package models

import (
	"github.com/erdemtoraman/opening-hours/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func TestTimeline_ToOutputFormat(t *testing.T) {
	var openingHoursInput OpeningHours
	inputFiles := make(map[string]string)
	outputFiles := make(map[string]string)
	err := filepath.Walk("../examples", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".in" {
			inputFiles[strings.TrimSuffix(info.Name(), ".in")] = path
		}
		if filepath.Ext(path) == ".out" {
			outputFiles[strings.TrimSuffix(info.Name(), ".out")] = path
		}
		return nil
	})
	assert.NoError(t, err)
	for key, filepath := range inputFiles {
		log.Println(filepath)
		err := helper.ReadFromFileAndPopulateDTO(filepath, &openingHoursInput)
		assert.NoError(t, err)
		expectedResult, err := ioutil.ReadFile(outputFiles[key])
		assert.NoError(t, err)
		timeline, err := openingHoursInput.ConvertToTimeline()
		assert.NoError(t, err)
		result := timeline.ToOutputFormat().String()
		assert.Equal(t, string(expectedResult), result)
	}

}
