package cmd

import (
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

func newReposCmd() *cli.Command {
	return &cli.Command{
		Name:   "repos",
		Usage:  "Gets repos stats (number of repos, forks, pulls)",
		Action: cmdRepos,
		Flags:  flags(),
	}
}

func cmdRepos(cCtx *cli.Context) error {
	gh, err := initCmd(cCtx)
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{RepoStats: &stats.GitHubRepoStats{}}
	stats.RepoStats.NumRepos, err = gh.NumRepos()
	if err != nil {
		return err
	}

	stats.RepoStats.NumForks, err = gh.NumForks()
	if err != nil {
		return err
	}

	stats.PullStats, err = gh.NumPulls()
	if err != nil {
		return err
	}

	return writeStats(cCtx, gh, stats)
}
