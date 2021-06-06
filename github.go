package extra

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xwjdsh/awesome-go-extra/models"
)

type GitHubAPI interface {
	UnmarshalGitHubRepo(ctx context.Context, fullName string, r *models.Record) error
}

var _ GitHubAPI = new(GitHubClient)

type GitHubClient struct {
	Username string
	Token    string
}

func NewGitHubClient(username, token string) *GitHubClient {
	return &GitHubClient{
		Username: username,
		Token:    token,
	}
}

func (c *GitHubClient) UnmarshalGitHubRepo(ctx context.Context, fullName string, r *models.Record) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.github.com/repos/%s", fullName), nil)
	if err != nil {
		return err
	}
	if c.Token != "" {
		req.SetBasicAuth(c.Username, c.Token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, r)
}
