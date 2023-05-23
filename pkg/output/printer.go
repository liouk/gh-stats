package output

import (
	"fmt"
	"io"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/stats"
)

func Print(writer io.Writer, stats *stats.GitHubViewerStats) {
	if stats.RepoStats != nil {
		fmt.Fprintf(writer, "%sRepos: %d\n", icons.Repo, stats.RepoStats.NumRepos)
		fmt.Fprintf(writer, "%sForks: %d\n", icons.Fork, stats.RepoStats.NumForks)
		fmt.Fprintf(writer, "%sPulls: %d\n", icons.Pull, stats.RepoStats.NumPulls)
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
}
