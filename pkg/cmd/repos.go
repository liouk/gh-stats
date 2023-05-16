package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/urfave/cli/v2"
)

func newReposCmd() *cli.Command {
	return &cli.Command{
		Name:   "repos",
		Usage:  "Gets repos stats (number of repos, forks, pulls)",
		Action: cmdRepos,
	}
}

func cmdRepos(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	return cmdReposWithGitHubContext(cCtx, gh)
}

func cmdReposWithGitHubContext(_ *cli.Context, gh *github.AuthenticatedGitHubContext) error {
	numRepos, err := gh.NumRepos()
	if err != nil {
		return err
	}
	log.Logf("%sRepos: %d\n", icons.Repo, numRepos)

	numForks, err := gh.NumForks()
	if err != nil {
		return err
	}
	log.Logf("%sForks: %d\n", icons.Fork, numForks)

	numPulls, err := gh.NumPulls()
	if err != nil {
		return err
	}
	log.Logf("%sPulls: %d\n", icons.Pull, numPulls)

	return nil
}
