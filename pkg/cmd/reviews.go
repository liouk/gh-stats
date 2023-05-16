package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/urfave/cli/v2"
)

func newReviewsCmd() *cli.Command {
	return &cli.Command{
		Name:   "reviews",
		Usage:  "Gets reviews stats",
		Action: cmdReviews,
	}
}

func cmdReviews(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	return cmdReviewsWithGitHubContext(cCtx, gh)
}

func cmdReviewsWithGitHubContext(_ *cli.Context, gh *github.AuthenticatedGitHubContext) error {
	numReviews, err := gh.NumReviews()
	if err != nil {
		return err
	}

	log.Logf("%sReviews: %d\n", icons.Review, numReviews)
	return nil
}
