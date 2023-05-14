package cmd

import (
	"context"
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
	ctx := context.Background()
	ghClient := github.NewAuthenticatedClient()
	numRepos, err := github.NumRepos(ctx, ghClient)
	if err != nil {
		return err
	}
	fmt.Println("repos:", numRepos)

	numForks, err := github.NumForks(ctx, ghClient)
	if err != nil {
		return err
	}
	fmt.Println("forks:", numForks)

	numPulls, err := github.NumPulls(ctx, ghClient)
	if err != nil {
		return err
	}
	fmt.Println("pulls:", numPulls)

	return nil
}
