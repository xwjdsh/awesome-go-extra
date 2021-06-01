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
)

var (
	cacheDuration  = flag.Duration("cd", 24*time.Hour, "cache duration")
	githubUsername = os.Getenv("EXTRA_GITHUB_USERNAME")
	githubToken    = os.Getenv("EXTRA_GITHUB_TOKEN")
)

func main() {
	flag.Parse()

	h := extra.New(*cacheDuration, githubUsername, githubToken)
	r, err := h.GetResult(context.Background())
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	buf.WriteString("# Awesome Go Extra\n")
	for _, c := range r.Categories {
		buf.WriteString(fmt.Sprintf("%s %s\n", c.Heading.ToMD(), c.Text))
		if c.Desc != "" {
			buf.WriteString(fmt.Sprintf("*%s*\n", c.Desc))
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
					r.Description,
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
