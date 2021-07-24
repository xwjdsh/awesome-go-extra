package main

import (
	"context"
	"flag"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	extra "github.com/xwjdsh/awesome-go-extra"
	"github.com/xwjdsh/awesome-go-extra/models"
)

var (
	refreshCache   = flag.Bool("refresh-cache", false, "refresh cache")
	tmplFilePath   = flag.String("tmpl", "", "template file path")
	order          = flag.String("order", "star", "order by")
	outputFilePath = flag.String("w", "", "output file path")
	githubUsername = os.Getenv("EXTRA_GITHUB_USERNAME")
	githubToken    = os.Getenv("EXTRA_GITHUB_TOKEN")
)

func main() {
	flag.Parse()

	cacheFilePath := "./repos.db"
	modelsHandler, err := models.Init(cacheFilePath)
	if err != nil {
		panic(err)
	}

	githubClient := extra.NewGitHubClient(githubUsername, githubToken)
	h := extra.New(*refreshCache, cacheFilePath, "./mapping.json", "./awesome-go/README.md", githubClient, modelsHandler)
	cas, err := h.GetResult(context.Background())
	if err != nil {
		panic(err)
	}

	sortByStar := func(record1, record2 *models.Record) bool {
		if record1.Repo.StargazersCount == record2.Repo.StargazersCount {
			if record1.Repo.ForksCount == record2.Repo.ForksCount {
				return record1.Repo.OpenIssuesCount < record2.Repo.OpenIssuesCount
			}
			return record1.Repo.ForksCount > record2.Repo.ForksCount
		}
		return record1.Repo.StargazersCount > record2.Repo.StargazersCount
	}

	for _, ca := range cas {
		sort.Slice(ca.Records, func(i, j int) bool {
			record1, record2 := ca.Records[i], ca.Records[j]
			switch *order {
			case "star":
				return sortByStar(record1, record2)
			case "pushed":
				if record1.Repo.PushedAt.Equal(record2.Repo.PushedAt) {
					return sortByStar(record1, record2)
				}
				return record1.Repo.PushedAt.After(record2.Repo.PushedAt)
			case "created":
				if record1.Repo.CreatedAt.Equal(record2.Repo.CreatedAt) {
					return sortByStar(record1, record2)
				}
				return record1.Repo.CreatedAt.Before(record2.Repo.CreatedAt)
			default:
				return sortByStar(record1, record2)
			}
		})
	}

	t, err := template.New(filepath.Base(*tmplFilePath)).Funcs(template.FuncMap{
		"headingMD": func(heading models.Heading) string {
			return heading.ToMD()
		},
		"ifelse": func(b bool, i, j interface{}) interface{} {
			if b {
				return i
			}
			return j
		},
		"recordAttr": func(isGithubRepo bool, attr interface{}) string {
			if !isGithubRepo {
				return "-"
			}
			switch v := attr.(type) {
			case string:
				return v
			case int:
				return strconv.Itoa(v)
			case time.Time:
				return v.Format(time.RFC3339)
			}
			return ""
		},
		"replace": strings.Replace,
	}).ParseFiles(*tmplFilePath)
	if err != nil {
		panic(err)
	}

	writer := os.Stdout
	if *outputFilePath != "" {
		f, err := os.OpenFile(*outputFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		writer = f
	}

	result := struct {
		Categories []*models.Category
		OrderBy    string
	}{
		Categories: cas,
		OrderBy:    *order,
	}

	if err := t.Execute(writer, result); err != nil {
		panic(err)
	}
}
