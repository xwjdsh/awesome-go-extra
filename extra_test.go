package extra

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	h := &Handler{
		ignoreGitHubRequest: true,
	}
	data, err := ioutil.ReadFile("./awesome-go/README.md")
	if err != nil {
		t.Error(err)
	}
	categories, err := h.parse(context.Background(), data)
	if err != nil {
		t.Error(err)
	}
	if len(categories) == 0 {
		t.Error("no category")
	}
}
