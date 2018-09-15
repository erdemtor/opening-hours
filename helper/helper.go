package helper

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFromFileAndPopulateDTO(fileAddress string, data interface{}) error {
	fileContent, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileContent, data)
}
