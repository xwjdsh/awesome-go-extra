package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	extra "github.com/xwjdsh/awesome-go-extra"
)

func main() {
	h := extra.New()
	h.SetBasicAuth(os.Getenv("EXTRA_GITHUB_USERNAME"), os.Getenv("EXTRA_GITHUB_TOKEN"))

	cas, err := h.GetResult(context.Background())
	if err != nil {
		panic(err)
	}
	println(len(cas))

	var buf strings.Builder
	buf.WriteString("# Awesome Go Extra\n")
	for _, c := range cas {
		buf.WriteString(fmt.Sprintf("## %s\n", c.Text))
		if c.Desc != "" {
			buf.WriteString(fmt.Sprintf("*%s*", c.Desc))
		}
		if len(c.Records) == 0 {
			continue
		}

		buf.WriteString("|Name|Desc|Star|PushedAt|OpenIssues|\n")
		buf.WriteString(strings.Repeat("|:---:", 5) + "|\n")
		for _, r := range c.Records {
			nameLink := fmt.Sprintf("[%s](%s)", r.Name, r.URL)
			if r.Archived {
				nameLink = "ARCHIVED " + nameLink

			}
			buf.WriteString(
				fmt.Sprintf(
					"|%s|%s|%d|%s|%d|\n",
					nameLink,
					r.Description,
					r.StargazersCount,
					r.PushedAt.Format(time.RFC3339),
					r.OpenIssuesCount,
				),
			)
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf.String())
}
