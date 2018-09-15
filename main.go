package main

import (
	"flag"
	"github.com/erdemtoraman/opening-hours/helper"
	"github.com/erdemtoraman/opening-hours/models"
	"log"
)

var filepath string

func init() {
	flag.StringVar(&filepath, "file", "", "please provide the filepath of the input file")

}
func main() {
	flag.Parse()
	input := models.OpeningHours{}
	err := helper.ReadFromFileAndPopulateDTO(filepath, &input)
	if err != nil {
		log.Fatalln(err)
	}
	timeline, err := input.ToTimeline()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(timeline.ToOutputFormat().String())
}
