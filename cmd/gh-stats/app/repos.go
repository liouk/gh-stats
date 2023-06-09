package app

import (
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

var cmdRepos = &cli.Command{
	Name:   "repos",
	Usage:  "Gets repos stats (number of repos, forks, pulls)",
	Action: runRepos,
	Flags:  flags(),
}

func runRepos(cCtx *cli.Context) error {
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
