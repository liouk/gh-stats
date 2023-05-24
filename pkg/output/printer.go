package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/stats"
)

func Print(writer io.Writer, stats *stats.GitHubViewerStats, outputType string, withIcons bool) error {
	if strings.EqualFold(outputType, "json") {
		bytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			return err
		}

		fmt.Fprintf(writer, string(bytes))
		return nil
	}

	var iconRepo, iconFork, iconPull, iconCommit, iconReview, iconCode string
	var langIcons map[string]string
	if withIcons {
		iconRepo = icons.Repo
		iconFork = icons.Fork
		iconPull = icons.Pull
		iconCommit = icons.Commit
		iconReview = icons.Review
		iconCode = icons.Code
		langIcons = icons.LangIcons
	}

	if stats.RepoStats != nil {
		fmt.Fprintf(writer, "%sRepos: %d\n", iconRepo, stats.RepoStats.NumRepos)
		fmt.Fprintf(writer, "%sForks: %d\n", iconFork, stats.RepoStats.NumForks)
		fmt.Fprintf(writer, "%sPulls: %d (%d open; %d closed; %d merged)\n",
			iconPull,
			stats.PullStats.TotalCount,
			stats.PullStats.OpenCount,
			stats.PullStats.ClosedCount,
			stats.PullStats.MergedCount,
		)
	}

	if stats.CommitStats != nil {
		fmt.Fprintf(writer, "%sCommits: %d\n", iconCommit, stats.CommitStats.NumCommits)
	}

	if stats.ReviewStats != nil {
		fmt.Fprintf(writer, "%sReviews: %d\n", iconReview, stats.ReviewStats.NumReviews)
	}

	if stats.LangStats != nil {
		fmt.Fprintf(writer, "%sLanguage stats:\n", iconCode)
		for _, lang := range stats.LangStats.Languages {
			fmt.Fprintf(writer, "   %s%s: %.2f%%\n", langIcons[lang.Name], lang.Name, lang.Perc)
		}
	}

	return nil
}
