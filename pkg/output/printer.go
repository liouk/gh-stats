package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/stats"
)

func Print(writer io.Writer, stats *stats.GitHubViewerStats, outputType string) error {
	if strings.EqualFold(outputType, "json") {
		bytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			return err
		}

		fmt.Fprintf(writer, string(bytes))
		return nil
	}

	if stats.RepoStats != nil {
		fmt.Fprintf(writer, "%sRepos: %d\n", icons.Repo, stats.RepoStats.NumRepos)
		fmt.Fprintf(writer, "%sForks: %d\n", icons.Fork, stats.RepoStats.NumForks)
		fmt.Fprintf(writer, "%sPulls: %d (%d open; %d closed; %d merged)\n",
			icons.Pull,
			stats.PullStats.TotalCount,
			stats.PullStats.OpenCount,
			stats.PullStats.ClosedCount,
			stats.PullStats.MergedCount,
		)
	}

	if stats.CommitStats != nil {
		fmt.Fprintf(writer, "%sCommits: %d\n", icons.Commit, stats.CommitStats.NumCommits)
	}

	if stats.ReviewStats != nil {
		fmt.Fprintf(writer, "%sReviews: %d\n", icons.Review, stats.ReviewStats.NumReviews)
	}

	if stats.LangStats != nil {
		fmt.Fprintf(writer, "%sLanguage stats:\n", icons.Code)
		for _, lang := range stats.LangStats.Languages {
			fmt.Fprintf(writer, "   %s %s: %.2f%%\n", icons.LangIcons[lang.Name], lang.Name, lang.Perc)
		}
	}

	return nil
}
