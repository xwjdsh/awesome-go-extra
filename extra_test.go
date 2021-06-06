package extra

import (
	"context"
	"testing"

	"github.com/xwjdsh/awesome-go-extra/models"
)

type MockModelsHandler struct{}

func (h *MockModelsHandler) SyncCategories(cas []*models.Category) error {
	return nil
}

func (h *MockModelsHandler) GetCategories() ([]*models.Category, error) {
	return []*models.Category{}, nil
}

type MockGitHubClient struct{}

func (c *MockGitHubClient) UnmarshalGitHubRepo(ctx context.Context, fullName string, r *models.Record) error {
	return nil
}

func TestGetResult(t *testing.T) {
	h := &Handler{
		mappingPth:          "./testdata/mapping.json",
		awesomeGoReadmePath: "./testdata/README.md",
		cachePath:           "./testdata/testrepos.db",
		modelsHandler:       new(MockModelsHandler),
		githubAPI:           new(MockGitHubClient),
	}

	categories, err := h.GetResult(context.Background())
	if err != nil {
		t.Error(err)
	}
	if len(categories) == 0 {
		t.Error("no category")
	}

	found := false
	for _, ca := range categories {
		for _, r := range ca.Records {
			if r.FullName == "go-joe/joe" {
				found = true
				break
			}
		}

		if found {
			break
		}
	}
	if !found {
		t.Error("mapping error")
	}
}
