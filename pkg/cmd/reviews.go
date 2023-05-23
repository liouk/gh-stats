package cmd

import (
	"os"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/liouk/gh-stats/pkg/output"
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

func newReviewsCmd() *cli.Command {
	return &cli.Command{
		Name:   "reviews",
		Usage:  "Gets reviews stats",
		Action: cmdReviews,
		Flags:  flags(),
	}
}

func cmdReviews(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{ReviewStats: &stats.GitHubReviewStats{}}
	stats.ReviewStats.NumReviews, err = gh.NumReviews()
	if err != nil {
		return err
	}

	output.Print(os.Stdout, stats)
	return nil
}
