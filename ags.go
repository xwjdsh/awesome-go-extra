package ags

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
)

const (
	AwesomeGoREADME = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
)

type Category struct {
	ID    string
	Text  string
	Desc  string
	Repos []*Repo
}

type Repo struct {
	ID        string
	Name      string
	Link      string
	Star      int
	UpdatedAt time.Time
}

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) GetResult(ctx context.Context) {

}

func (h *Handler) fetch(ctx context.Context) ([]byte, error) {
	resp, err := http.Get(AwesomeGoREADME)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (h *Handler) parse(data []byte) error {
	html := markdown.ToHTML(data, nil, nil)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return err
	}

	prefixMap := map[string]string{}
	// categories := []*Category{}
	doc.Find(`li a`).Each(func(i int, s *goquery.Selection) {
		if id := s.AttrOr("href", "#"); strings.HasPrefix(id, "#") {
			id = id[1:]
			prefix := ""
			p := s.Parent().Parent()
			if goquery.NodeName(p) == "ul" {
				pp := p.Parent()
				if goquery.NodeName(pp) == "li" {
					prefix = pp.Find("a").AttrOr("href", "#")[1:]
					if prefix == "awesome-go" {
						prefix = ""
					}
				}
			}

			newID := id
			if prefix != "" {
				if v, ok := prefixMap[prefix]; ok {
					prefix = v
				}
				newID = prefix + ">" + id
			}
			prefixMap[id] = newID
			println(newID)
		}
	})
	return nil
}
