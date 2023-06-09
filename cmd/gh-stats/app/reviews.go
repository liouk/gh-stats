package app

import (
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

var cmdReviews = &cli.Command{
	Name:   "reviews",
	Usage:  "Gets reviews stats",
	Action: runReviews,
	Flags:  flags(),
}

func runReviews(cCtx *cli.Context) error {
	gh, err := initCmd(cCtx)
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{ReviewStats: &stats.GitHubReviewStats{}}
	stats.ReviewStats.NumReviews, err = gh.NumReviews()
	if err != nil {
		return err
	}

	return writeStats(cCtx, gh, stats)
}
