package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Handler the db Handler.
type Handler struct {
	db *gorm.DB
}

// Category is the sort of repos
type Category struct {
	ID           int64   `gorm:"primarykey"`
	Text         string  `gorm:"column:text;unique"`
	HeadingLevel Heading `gorm:"column:heading_level"`
	Desc         string  `gorm:"column:desc"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Records      []*Record
}

// TableName ...
func (Category) TableName() string {
	return "category"
}

// Record record on awesome-go
type Record struct {
	ID           int64      `gorm:"primarykey"`
	Name         string     `gorm:"column:name"`
	FullName     string     `gorm:"column:full_name"`
	Description  string     `gorm:"column:description"`
	URL          string     `gorm:"column:url;unique"`
	Repo         GitHubRepo `gorm:"column:repo"`
	IsGitHubRepo bool       `gorm:"column:is_github_repo"`
	CategoryID   int64      `gorm:"column:category_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
		now := time.Now()
		for _, ca := range cas {
			ca.UpdatedAt = now
			if err := tx.Omit("Records").Clauses(clause.OnConflict{
				UpdateAll: true,
				Columns:   []clause.Column{{Name: "text"}},
			}).Create(ca).Error; err != nil {
				return err
			}
			ca1 := &Category{}
			if err := tx.Select("id").First(ca1, "text = ?", ca.Text).Error; err != nil {
				return err
			}
			for _, r := range ca.Records {
				r.CategoryID = ca1.ID
				r.UpdatedAt = now
				if err := tx.Clauses(clause.OnConflict{
					UpdateAll: true,
					Columns:   []clause.Column{{Name: "url"}},
				}).Create(r).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Where("updated_at != ?", now).Delete(&Category{}).Error; err != nil {
			return err
		}
		if err := tx.Where("updated_at != ?", now).Delete(&Record{}).Error; err != nil {
			return err
		}

		return nil
	})
}

type GitHubRepo struct {
	ID       int    `json:"id,omitempty"`
	NodeID   string `json:"node_id,omitempty"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Private  bool   `json:"private,omitempty"`
	Owner    struct {
		Login             string `json:"login,omitempty"`
		ID                int    `json:"id,omitempty"`
		NodeID            string `json:"node_id,omitempty"`
		AvatarURL         string `json:"avatar_url,omitempty"`
		GravatarID        string `json:"gravatar_id,omitempty"`
		URL               string `json:"url,omitempty"`
		HTMLURL           string `json:"html_url,omitempty"`
		FollowersURL      string `json:"followers_url,omitempty"`
		FollowingURL      string `json:"following_url,omitempty"`
		GistsURL          string `json:"gists_url,omitempty"`
		StarredURL        string `json:"starred_url,omitempty"`
		SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
		OrganizationsURL  string `json:"organizations_url,omitempty"`
		ReposURL          string `json:"repos_url,omitempty"`
		EventsURL         string `json:"events_url,omitempty"`
		ReceivedEventsURL string `json:"received_events_url,omitempty"`
		Type              string `json:"type,omitempty"`
		SiteAdmin         bool   `json:"site_admin,omitempty"`
	} `json:"owner,omitempty"`
	HTMLURL          string      `json:"html_url,omitempty"`
	Description      string      `json:"description,omitempty"`
	Fork             bool        `json:"fork,omitempty"`
	URL              string      `json:"url,omitempty"`
	ForksURL         string      `json:"forks_url,omitempty"`
	KeysURL          string      `json:"keys_url,omitempty"`
	CollaboratorsURL string      `json:"collaborators_url,omitempty"`
	TeamsURL         string      `json:"teams_url,omitempty"`
	HooksURL         string      `json:"hooks_url,omitempty"`
	IssueEventsURL   string      `json:"issue_events_url,omitempty"`
	EventsURL        string      `json:"events_url,omitempty"`
	AssigneesURL     string      `json:"assignees_url,omitempty"`
	BranchesURL      string      `json:"branches_url,omitempty"`
	TagsURL          string      `json:"tags_url,omitempty"`
	BlobsURL         string      `json:"blobs_url,omitempty"`
	GitTagsURL       string      `json:"git_tags_url,omitempty"`
	GitRefsURL       string      `json:"git_refs_url,omitempty"`
	TreesURL         string      `json:"trees_url,omitempty"`
	StatusesURL      string      `json:"statuses_url,omitempty"`
	LanguagesURL     string      `json:"languages_url,omitempty"`
	StargazersURL    string      `json:"stargazers_url,omitempty"`
	ContributorsURL  string      `json:"contributors_url,omitempty"`
	SubscribersURL   string      `json:"subscribers_url,omitempty"`
	SubscriptionURL  string      `json:"subscription_url,omitempty"`
	CommitsURL       string      `json:"commits_url,omitempty"`
	GitCommitsURL    string      `json:"git_commits_url,omitempty"`
	CommentsURL      string      `json:"comments_url,omitempty"`
	IssueCommentURL  string      `json:"issue_comment_url,omitempty"`
	ContentsURL      string      `json:"contents_url,omitempty"`
	CompareURL       string      `json:"compare_url,omitempty"`
	MergesURL        string      `json:"merges_url,omitempty"`
	ArchiveURL       string      `json:"archive_url,omitempty"`
	DownloadsURL     string      `json:"downloads_url,omitempty"`
	IssuesURL        string      `json:"issues_url,omitempty"`
	PullsURL         string      `json:"pulls_url,omitempty"`
	MilestonesURL    string      `json:"milestones_url,omitempty"`
	NotificationsURL string      `json:"notifications_url,omitempty"`
	LabelsURL        string      `json:"labels_url,omitempty"`
	ReleasesURL      string      `json:"releases_url,omitempty"`
	DeploymentsURL   string      `json:"deployments_url,omitempty"`
	CreatedAt        time.Time   `json:"created_at,omitempty"`
	UpdatedAt        time.Time   `json:"updated_at,omitempty"`
	PushedAt         time.Time   `json:"pushed_at,omitempty"`
	GitURL           string      `json:"git_url,omitempty"`
	SSHURL           string      `json:"ssh_url,omitempty"`
	CloneURL         string      `json:"clone_url,omitempty"`
	SvnURL           string      `json:"svn_url,omitempty"`
	Homepage         interface{} `json:"homepage,omitempty"`
	Size             int         `json:"size,omitempty"`
	StargazersCount  int         `json:"stargazers_count,omitempty"`
	WatchersCount    int         `json:"watchers_count,omitempty"`
	Language         string      `json:"language,omitempty"`
	HasIssues        bool        `json:"has_issues,omitempty"`
	HasProjects      bool        `json:"has_projects,omitempty"`
	HasDownloads     bool        `json:"has_downloads,omitempty"`
	HasWiki          bool        `json:"has_wiki,omitempty"`
	HasPages         bool        `json:"has_pages,omitempty"`
	ForksCount       int         `json:"forks_count,omitempty"`
	MirrorURL        interface{} `json:"mirror_url,omitempty"`
	Archived         bool        `json:"archived,omitempty"`
	Disabled         bool        `json:"disabled,omitempty"`
	OpenIssuesCount  int         `json:"open_issues_count,omitempty"`
	License          struct {
		Key    string `json:"key,omitempty"`
		Name   string `json:"name,omitempty"`
		SpdxID string `json:"spdx_id,omitempty"`
		URL    string `json:"url,omitempty"`
		NodeID string `json:"node_id,omitempty"`
	} `json:"license,omitempty"`
	Forks            int         `json:"forks,omitempty"`
	OpenIssues       int         `json:"open_issues,omitempty"`
	Watchers         int         `json:"watchers,omitempty"`
	DefaultBranch    string      `json:"default_branch,omitempty"`
	TempCloneToken   interface{} `json:"temp_clone_token,omitempty"`
	NetworkCount     int         `json:"network_count,omitempty"`
	SubscribersCount int         `json:"subscribers_count,omitempty"`
}

func (r *GitHubRepo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal value:", value))
	}

	return json.Unmarshal(data, r)
}

func (r GitHubRepo) Value() (driver.Value, error) {
	if r.ID == 0 {
		return nil, nil
	}
	return json.Marshal(r)
}
