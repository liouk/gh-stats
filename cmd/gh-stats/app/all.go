package app

import (
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/urfave/cli/v2"
)

var cmdAll = &cli.Command{
	Name:   "all",
	Usage:  "Gets all stats (repos, forks, pulls, commits, reviews, languages)",
	Action: runAll,
	Flags: flags(
		flagLangNum,
		flagLangIgnore,
	),
}

func runAll(cCtx *cli.Context) error {
	gh, err := initCmd(cCtx)
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{
		RepoStats:   &stats.GitHubRepoStats{},
		CommitStats: &stats.GitHubCommitStats{},
		ReviewStats: &stats.GitHubReviewStats{},
		LangStats:   &stats.GitHubLangStats{},
	}

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

	stats.CommitStats.NumCommits, err = gh.NumCommits()
	if err != nil {
		return err
	}

	stats.ReviewStats.NumReviews, err = gh.NumReviews()
	if err != nil {
		return err
	}

	stats.LangStats.Languages, err = gh.LangStats(cCtx.Int("num"), cCtx.StringSlice("ignore"))
	if err != nil {
		return err
	}

	return writeStats(cCtx, gh, stats)
}
