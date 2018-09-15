package main

import (
	"./helper"
	"./models"
	"flag"
	"fmt"
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
