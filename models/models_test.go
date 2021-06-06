package models

import (
	"os"
	"testing"
)

func TestGetCategorise(t *testing.T) {
	dbPath := "./test_models.db"
	h, err := Init(dbPath)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dbPath)

	_, err = h.GetCategories()
	if err != nil {
		t.Error(err)
	}
}

func TestSyncCategorise(t *testing.T) {
	dbPath := "./test_models.db"
	h, err := Init(dbPath)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dbPath)

	cas := []*Category{
		{
			Text: "test",
			Records: []*Record{
				{
					Name: "test",
				},
			},
		},
	}
	if err = h.SyncCategories(cas); err != nil {
		t.Error(err)
	}

	categories, err := h.GetCategories()
	if err != nil {
		t.Error(err)
	}
	if len(categories) != 1 || categories[0].Text != "test" {
		t.Error("unexpected categories")
	}
	ca := categories[0]
	if len(ca.Records) != 1 || ca.Records[0].Name != "test" {
		t.Error("unexpected records")
	}
}
