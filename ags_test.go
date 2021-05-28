package ags

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	h := New()
	data, err := ioutil.ReadFile("./testdata/README.md")
	if err != nil {
		t.Error(err)
	}
	h.parse(data)
}
