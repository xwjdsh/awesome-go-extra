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
		"sort": func(records []*models.Record) []*models.Record {
			sort.Slice(records, func(i, j int) bool {
				if records[i].Repo.StargazersCount == records[j].Repo.StargazersCount {
					if records[i].Repo.ForksCount == records[j].Repo.ForksCount {
						return records[i].Repo.OpenIssuesCount < records[j].Repo.OpenIssuesCount
					}
					return records[i].Repo.ForksCount > records[j].Repo.ForksCount
				}
				return records[i].Repo.StargazersCount > records[j].Repo.StargazersCount
			})
			return records
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

	if err := t.Execute(writer, cas); err != nil {
		panic(err)
	}
}
