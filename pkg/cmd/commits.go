package cmd

import (
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
	gh, err := initCmd(cCtx)
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{CommitStats: &stats.GitHubCommitStats{}}
	stats.CommitStats.NumCommits, err = gh.NumCommits()
	if err != nil {
		return err
	}

	return writeStats(cCtx, stats)
}
