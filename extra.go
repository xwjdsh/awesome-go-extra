package extra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

const (
	AwesomeGoREADME = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
)

type Category struct {
	ID      string
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
}

type Handler struct {
	GitHubAuthUsername string
	GitHubAuthToken    string
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetBasicAuth(username, token string) {
	h.GitHubAuthUsername = username
	h.GitHubAuthToken = token
}

func (h *Handler) GetResult(ctx context.Context) ([]*Category, error) {
	data, err := h.fetch(ctx)
	if err != nil {
		return nil, err
	}

	return h.parse(ctx, data)
}

func (h *Handler) fetch(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", AwesomeGoREADME, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
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
		next := doc.Find(id).Next()
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

		category := &Category{
			ID:   id,
			Text: text,
			Desc: desc,
		}
		if listSelection != nil {
			for i := range listSelection.Children().Nodes {
				s := listSelection.Children().Eq(i)

				a := s.Find("a")
				addr := a.AttrOr("href", "")
				name := a.Text()

				a.Remove()
				desc := strings.TrimPrefix(s.Text(), " - ")

				record := &Record{
					Name:        name,
					URL:         addr,
					Description: desc,
				}
				if prefix := "https://github.com/"; strings.HasPrefix(addr, prefix) {
					fullName := strings.TrimPrefix(addr, prefix)
					if err := h.getGitHubRepo(ctx, fullName, record); err != nil {
						log.Println(err.Error())
					}
				}
				category.Records = append(category.Records, record)
			}
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (h *Handler) getGitHubRepo(ctx context.Context, fullName string, r *Record) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.github.com/repos/%s", fullName), nil)
	if err != nil {
		return err
	}
	if h.GitHubAuthToken != "" {
		req.SetBasicAuth(h.GitHubAuthUsername, h.GitHubAuthToken)
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
