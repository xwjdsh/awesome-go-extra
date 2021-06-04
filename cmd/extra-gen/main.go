package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	extra "github.com/xwjdsh/awesome-go-extra"
	"github.com/xwjdsh/awesome-go-extra/models"
)

var (
	cacheFilePath  = flag.String("cache", "repos.db", "cache file path")
	githubUsername = os.Getenv("EXTRA_GITHUB_USERNAME")
	githubToken    = os.Getenv("EXTRA_GITHUB_TOKEN")
)

func main() {
	flag.Parse()

	modelsHandler, err := models.Init(*cacheFilePath)
	if err != nil {
		panic(err)
	}

	h := extra.New(*cacheFilePath, githubUsername, githubToken, modelsHandler)
	cas, err := h.GetResult(context.Background())
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	buf.WriteString("# Awesome Go Extra\n")
	for _, c := range cas {
		buf.WriteString(fmt.Sprintf("%s %s\n", c.HeadingLevel.ToMD(), c.Text))
		if c.Desc != "" {
			buf.WriteString(fmt.Sprintf("*%s*\n", c.Desc))
		}

		if len(c.Records) == 0 {
			continue
		}

		buf.WriteString("|Name|Description|Star|Open Issues|CreatedAt|PushedAt|\n")
		buf.WriteString(strings.Repeat("|:---:", 6) + "|\n")
		sort.Slice(c.Records, func(i, j int) bool {
			return c.Records[i].StargazersCount > c.Records[j].StargazersCount
		})
		for _, r := range c.Records {
			nameLink := fmt.Sprintf("[%s](%s)", r.Name, r.URL)
			if r.Archived {
				nameLink = "**[ARCHIVED]**  " + nameLink
			}
			buf.WriteString(
				fmt.Sprintf(
					"|%s|%s|%s|%s|%s|%s|\n",
					nameLink,
					strings.ReplaceAll(r.Description, "|", "`|`"),
					getRecordAttr(r.IsGitHubRepo, func() string { return strconv.Itoa((r.StargazersCount)) }),
					getRecordAttr(r.IsGitHubRepo, func() string { return strconv.Itoa((r.OpenIssuesCount)) }),
					getRecordAttr(r.IsGitHubRepo, func() string { return r.CreatedAt.Format(time.RFC3339) }),
					getRecordAttr(r.IsGitHubRepo, func() string { return r.PushedAt.Format(time.RFC3339) }),
				),
			)
		}
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
}

func getRecordAttr(isGithubRepo bool, f func() string) string {
	if !isGithubRepo {
		return "-"
	}

	return f()
}
