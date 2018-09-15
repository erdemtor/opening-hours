package main

import (
	"flag"
	"fmt"
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
		log.Println(err)
		return
	}
	timeline, err := input.ToTimeline()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(timeline.ToOutputFormat().String())
}
