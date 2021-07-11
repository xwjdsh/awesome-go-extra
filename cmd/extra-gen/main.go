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
	tmplFilePath   = flag.String("tmpl", "extra-md.tmpl", "template file path")
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
		"sort": func(records []*models.Record, orderBy string) []*models.Record {
			sort.Slice(records, func(i, j int) bool {
				switch orderBy {
				case "star":
					return records[i].StargazersCount > records[j].StargazersCount
				case "open_issues":
					return records[i].OpenIssuesCount < records[j].OpenIssuesCount
				case "pushed_at":
					return records[i].PushedAt.After(records[j].PushedAt)
				default:
					return records[i].StargazersCount > records[j].StargazersCount
				}
			})
			return records
		},
		"replace": strings.Replace,
	}).ParseFiles(*tmplFilePath)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, cas); err != nil {
		panic(err)
	}
}
