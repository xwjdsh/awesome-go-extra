package extra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

const (
	GitHubPrefix = "https://github.com/"
)

type Heading string

const (
	H1 Heading = "h1"
	H2 Heading = "h2"
	H3 Heading = "h3"
	H4 Heading = "h4"
	H5 Heading = "h5"
)

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

type Result struct {
	Time       time.Time   `json:"time"`
	Categories []*Category `json:"categories"`
}

type Category struct {
	ID      string
	Heading Heading
	Text    string
	Desc    string
	Records []*Record
}

type Record struct {
	Name            string    `json:"-"`
	FullName        string    `json:"full_name"`
	URL             string    `json:"html_url"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PushedAt        time.Time `json:"pushed_at"`
	StargazersCount int       `json:"stargazers_count"`
	Archived        bool      `json:"archived"`
	OpenIssuesCount int       `json:"open_issues_count"`
	IsGitHubRepo    bool      `json:"-"`
}

type Handler struct {
	ignoreGitHubRequest bool
	gitHubAuthUsername  string
	gitHubAuthToken     string
	cacheDuration       time.Duration
}

func New(cacheDuration time.Duration, gitHubAuthUsername, gitHubAuthToken string) *Handler {
	return &Handler{
		cacheDuration:      cacheDuration,
		gitHubAuthUsername: gitHubAuthUsername,
		gitHubAuthToken:    gitHubAuthToken,
	}
}

func (h *Handler) GetResult(ctx context.Context) (*Result, error) {
	result := &Result{}
	filePath := "./repo-data.json"
	if h.cacheDuration != 0 {
		if data, err := ioutil.ReadFile(filePath); err == nil {
			if err := json.Unmarshal(data, result); err == nil {
				if result.Time.Add(h.cacheDuration).After(time.Now()) {
					return result, nil
				}
			} else {
				os.Remove(filePath)
			}
		}
	}

	data, err := ioutil.ReadFile("./awesome-go/README.md")
	if err != nil {
		return nil, err
	}

	categories, err := h.parse(ctx, data)
	if err != nil {
		return nil, err
	}

	result.Time = time.Now()
	result.Categories = categories

	data, err = json.Marshal(result)
	if err != nil {
		return nil, err
	}

	ioutil.WriteFile(filePath, data, 0644)
	return result, nil
}

func (h *Handler) parse(ctx context.Context, data []byte) ([]*Category, error) {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.TOC
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(data, nil, renderer)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	categories := []*Category{}
	for i := 5; ; i++ {
		id := fmt.Sprintf("#toc_%d", i)
		text := doc.Find(fmt.Sprintf("a[href='%s']", id)).Text()
		if text == "" {
			break
		}

		desc := ""
		var listSelection *goquery.Selection
		headingSel := doc.Find(id)
		nodeName := goquery.NodeName(headingSel)
		next := headingSel.Next()
		switch goquery.NodeName(next) {
		case "p":
			desc = next.Text()
			nn := next.Next()
			if goquery.NodeName(nn) == "ul" {
				listSelection = nn
			}
		case "ul":
			listSelection = next
		}

		if listSelection != nil {
			category := &Category{
				ID:      id,
				Text:    text,
				Heading: Heading(nodeName),
				Desc:    desc,
			}
			categories = append(categories, category)
			for i := range listSelection.Children().Nodes {
				s := listSelection.Children().Eq(i)
				if sul := s.Find("ul"); sul.Length() > 0 {
					sul.Remove()
					listCategory := &Category{
						Text:    s.Text(),
						Heading: category.Heading.Sub(),
					}
					categories = append(categories, listCategory)
					for i := range sul.Children().Nodes {
						ss := sul.Children().Eq(i)
						if err := h.parseAndAddRecord(ctx, ss, listCategory); err != nil {
							return nil, err
						}
					}
					continue
				}
				if err := h.parseAndAddRecord(ctx, s, category); err != nil {
					return nil, err
				}
			}
		}
	}

	return categories, nil
}

func (h *Handler) getGitHubRepo(ctx context.Context, fullName string, r *Record) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.github.com/repos/%s", fullName), nil)
	if err != nil {
		return err
	}
	if h.gitHubAuthToken != "" {
		req.SetBasicAuth(h.gitHubAuthUsername, h.gitHubAuthToken)
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

func (h *Handler) parseAndAddRecord(ctx context.Context, s *goquery.Selection, category *Category) error {
	a := s.Find("a").First()
	addr := a.AttrOr("href", "")
	name := a.Text()

	a.Remove()
	desc := strings.TrimPrefix(s.Text(), " - ")

	record := &Record{
		Name:         name,
		FullName:     name,
		URL:          addr,
		Description:  desc,
		IsGitHubRepo: strings.HasPrefix(addr, GitHubPrefix),
	}
	if record.IsGitHubRepo && !h.ignoreGitHubRequest {
		fullName := strings.TrimPrefix(addr, GitHubPrefix)
		eles := strings.Split(fullName, "/")
		if len(eles) >= 2 {
			fullName = strings.Join(eles[:2], "/")
			if err := h.getGitHubRepo(ctx, fullName, record); err != nil {
				return fmt.Errorf("addr: %s, fullName: %s, err: %w", fullName, addr, err)
			}
		}
	}
	category.Records = append(category.Records, record)
	return nil
}
