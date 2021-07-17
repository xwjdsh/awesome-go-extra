package extra

import (
	"context"
	"testing"

	"github.com/xwjdsh/awesome-go-extra/models"
)

func TestUnmarshalGitHubRepo(t *testing.T) {
	client := NewGitHubClient("", "")
	r := new(models.GitHubRepo)
	err := client.UnmarshalGitHubRepo(context.Background(), "xwjdsh/awesome-go-extra", r)
	if err != nil {
		t.Error(err)
	}

	if r.URL == "" {
		t.Error("unmarshal error")
	}
}
