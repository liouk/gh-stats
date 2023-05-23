package cmd

import (
	"os"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/liouk/gh-stats/pkg/output"
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

func newCommitsCmd() *cli.Command {
	return &cli.Command{
		Name:   "commits",
		Usage:  "Gets commits stats",
		Action: cmdCommits,
		Flags:  flags(),
	}
}

func cmdCommits(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{CommitStats: &stats.GitHubCommitStats{}}
	stats.CommitStats.NumCommits, err = gh.NumCommits()
	if err != nil {
		return err
	}

	output.Print(os.Stdout, stats)
	return nil
}
