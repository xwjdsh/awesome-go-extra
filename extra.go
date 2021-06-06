package extra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"

	"github.com/xwjdsh/awesome-go-extra/models"
)

// Handler extra Handler
type Handler struct {
	modelsHandler       ModelsHandler
	githubAPI           GitHubAPI
	mapping             map[string]string
	cachePath           string
	mappingPth          string
	awesomeGoReadmePath string
}

type ModelsHandler interface {
	SyncCategories(cas []*models.Category) error
	GetCategories() ([]*models.Category, error)
}

// New returns a new Handler
func New(cachePath, mappingPth, awesomeGoReadmePath string, githubClient GitHubAPI, modelsHandler ModelsHandler) *Handler {
	h := &Handler{
		cachePath:           cachePath,
		mappingPth:          mappingPth,
		awesomeGoReadmePath: awesomeGoReadmePath,
		githubAPI:           githubClient,
		modelsHandler:       modelsHandler,
	}

	return h
}

// GetResult returns categories and records data
func (h *Handler) GetResult(ctx context.Context) ([]*models.Category, error) {
	if h.cachePath != "" {
		cas, err := h.modelsHandler.GetCategories()
		if err != nil {
			return nil, err
		}

		if len(cas) > 0 {
			return cas, nil
		}
	}

	data, err := ioutil.ReadFile(h.awesomeGoReadmePath)
	if err != nil {
		return nil, err
	}

	mappingData, err := ioutil.ReadFile(h.mappingPth)
	if err != nil {
		return nil, err
	}
	h.mapping = map[string]string{}
	if err := json.Unmarshal(mappingData, &h.mapping); err != nil {
		return nil, err
	}

	cas, err := h.parse(ctx, data)
	if err != nil {
		return nil, err
	}

	if h.cachePath != "" {
		if err := h.modelsHandler.SyncCategories(cas); err != nil {
			return nil, err
		}
	}

	return cas, nil
}

func (h *Handler) parse(ctx context.Context, data []byte) ([]*models.Category, error) {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.TOC
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(data, nil, renderer)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	categories := []*models.Category{}
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
			category := &models.Category{
				Text:         text,
				HeadingLevel: models.Heading(nodeName),
				Desc:         desc,
			}
			categories = append(categories, category)
			for i := range listSelection.Children().Nodes {
				s := listSelection.Children().Eq(i)
				if sul := s.Find("ul"); sul.Length() > 0 {
					sul.Remove()
					listCategory := &models.Category{
						Text:         s.Text(),
						HeadingLevel: category.HeadingLevel.Sub(),
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

func (h *Handler) parseAndAddRecord(ctx context.Context, s *goquery.Selection, category *models.Category) error {
	a := s.Find("a").First()
	addr := a.AttrOr("href", "")
	name := a.Text()

	if h.mapping != nil {
		if newAddr, ok := h.mapping[addr]; ok {
			addr = newAddr
		}
	}

	a.Remove()
	desc := strings.TrimPrefix(s.Text(), " - ")

	fullName, isGitHubRepo := getGitHubRepoFullName(addr)
	record := &models.Record{
		Name:         name,
		FullName:     fullName,
		URL:          addr,
		Description:  desc,
		IsGitHubRepo: isGitHubRepo,
	}
	if isGitHubRepo {
		if err := h.githubAPI.UnmarshalGitHubRepo(ctx, fullName, record); err != nil {
			return fmt.Errorf("addr: %s, fullName: %s, err: %w", fullName, addr, err)
		}
	}
	category.Records = append(category.Records, record)
	return nil
}

func getGitHubRepoFullName(addr string) (string, bool) {
	fullName := strings.TrimPrefix(addr, "https://github.com/")
	fullName = strings.TrimPrefix(fullName, "http://github.com/")
	if addr == fullName {
		return addr, false
	}

	elements := strings.Split(fullName, "/")
	if len(elements) < 2 {
		return addr, false
	}

	return strings.Join(elements[:2], "/"), true
}
