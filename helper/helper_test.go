package helper

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFromFileAndPopulateDTO(t *testing.T) {
	content := []byte(`{"data_key": "data_val"}`)
	tmpfile, err := ioutil.TempFile("", "*.in")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	json := map[string]string{}
	err = ReadFromFileAndPopulateDTO(tmpfile.Name(), &json)
	assert.NoError(t, err)
	assert.Contains(t, json, "data_key")
	assert.Equal(t, "data_val", json["data_key"])

	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	err = ReadFromFileAndPopulateDTO("nonexistentfile", &json)
	assert.Error(t, err)
}
