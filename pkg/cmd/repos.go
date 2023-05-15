package cmd

import (
	"fmt"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli/v2"
)

func newReposCmd() *cli.Command {
	return &cli.Command{
		Name:   "repos",
		Usage:  "Gets repos stats",
		Action: cmdRepos,
	}
}

func cmdRepos(cCtx *cli.Context) error {
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	numRepos, err := gh.NumRepos()
	if err != nil {
		return err
	}
	fmt.Println("repos:", numRepos)

	numForks, err := gh.NumForks()
	if err != nil {
		return err
	}
	fmt.Println("forks:", numForks)

	numPulls, err := gh.NumPulls()
	if err != nil {
		return err
	}
	fmt.Println("pulls:", numPulls)

	return nil
}
