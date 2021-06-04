package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Handler the db Handler.
type Handler struct {
	db *gorm.DB
}

// Category is the sort of repos
type Category struct {
	ID           int64   `gorm:"primarykey"`
	Text         string  `gorm:"column:text"`
	HeadingLevel Heading `gorm:"column:heading_level"`
	Desc         string  `gorm:"column:desc"`
	CreatedAt    time.Time
	Records      []*Record
}

// TableName ...
func (Category) TableName() string {
	return "category"
}

// Record record on awesome-go
type Record struct {
	ID              int64     `gorm:"primarykey"`
	Name            string    `gorm:"column:name"`
	FullName        string    `gorm:"column:full_name" json:"full_name"`
	URL             string    `gorm:"column:url" json:"html_url"`
	Description     string    `gorm:"column:description" json:"description"`
	RepoCreatedAt   time.Time `gorm:"column:repo_created_at" json:"created_at"`
	PushedAt        time.Time `gorm:"column:pushed_at" json:"pushed_at"`
	StargazersCount int       `gorm:"column:stargazers_count" json:"stargazers_count"`
	OpenIssuesCount int       `gorm:"column:open_issues_count" json:"open_issues_count"`
	Archived        bool      `gorm:"column:archived" json:"archived"`
	IsGitHubRepo    bool      `gorm:"column:is_github_repo"`
	CategoryID      int64     `gorm:"column:category_id"`
	CreatedAt       time.Time
}

// TableName ...
func (Record) TableName() string {
	return "record"
}

// Heading ...
type Heading string

const (
	// H1 ...
	H1 Heading = "h1"
	// H2 ...
	H2 Heading = "h2"
	// H3 ...
	H3 Heading = "h3"
	// H4 ...
	H4 Heading = "h4"
	// H5 ...
	H5 Heading = "h5"
)

// ToMD convert to markdown format
func (h Heading) ToMD() string {
	r := ""
	switch h {
	case H1:
		r = "#"
	case H2:
		r = "##"
	case H3:
		r = "###"
	case H4:
		r = "####"
	case H5:
		r = "#####"
	}

	return r
}

// Sub returns the next level of heading
func (h Heading) Sub() Heading {
	r := H5
	switch h {
	case H1:
		r = H2
	case H2:
		r = H3
	case H3:
		r = H4
	case H4:
		r = H5
	}

	return r
}

// Init returns a new Handler
func Init(dbPath string) (*Handler, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Category{}, &Record{}); err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

// GetCategories returns categories data.
func (h *Handler) GetCategories() ([]*Category, error) {
	cas := []*Category{}
	return cas, h.db.Preload("Records").Find(&cas).Error
}

// SyncCategories recreate categories and records.
func (h *Handler) SyncCategories(cas []*Category) error {
	return h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("1 = 1").Delete(&Category{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1 = 1").Delete(&Record{}).Error; err != nil {
			return err
		}

		for _, ca := range cas {
			if err := tx.Create(ca).Error; err != nil {
				return err
			}

			if err := tx.Model(&Category{ID: ca.ID}).Association("Records").Append(ca.Records); err != nil {
				return err
			}
		}

		return nil
	})
}
