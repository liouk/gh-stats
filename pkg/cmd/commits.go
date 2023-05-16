package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/urfave/cli/v2"
)

func newCommitsCmd() *cli.Command {
	return &cli.Command{
		Name:   "commits",
		Usage:  "Gets commits stats",
		Action: cmdCommits,
	}
}

func cmdCommits(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	return cmdCommitsWithGitHubContext(cCtx, gh)
}

func cmdCommitsWithGitHubContext(cCtx *cli.Context, gh *github.AuthenticatedGitHubContext) error {
	numCommits, err := gh.NumCommits()
	if err != nil {
		return err
	}

	log.Logf("%sCommits: %d\n", icons.Commit, numCommits)
	return nil
}
