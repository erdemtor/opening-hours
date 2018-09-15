package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	os.Args = []string{os.Args[0], "-file", "./examples/1.in"}
	outStandard := captureOutput(func() { main() })
	content, err := ioutil.ReadFile("./examples/1.out")
	assert.NoError(t, err)
	assert.EqualValues(t, strings.TrimSpace(string(content)), strings.TrimSpace(outStandard))

	os.Args = []string{os.Args[0], "-file", "nonexistent"}
	outStandard = captureOutput(func() { main() })
	assert.NoError(t, err)
	assert.Empty(t, outStandard)

}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
