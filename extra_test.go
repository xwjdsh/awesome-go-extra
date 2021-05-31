package extra

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	h := &Handler{
		GitHubAuthUsername: os.Getenv("TEST_EXTRA_USERNAME"),
		GitHubAuthToken:    os.Getenv("TEST_EXTRA_PASSWORD"),
	}
	data, err := ioutil.ReadFile("./testdata/README.md")
	if err != nil {
		t.Error(err)
	}
	categories, err := h.parse(data)
	if err != nil {
		t.Error(err)
	}
	if len(categories) == 0 {
		t.Error("no category")
	}
}
