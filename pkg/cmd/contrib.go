package cmd

import (
	"context"
	"fmt"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli/v2"
)

func newContribCmd() *cli.Command {
	return &cli.Command{
		Name:   "contrib",
		Usage:  "Gets contribution stats",
		Action: cmdContrib,
	}
}

func cmdContrib(cCtx *cli.Context) error {
	ctx := context.Background()
	ghClient := github.NewAuthenticatedClient()
	numCommits, err := github.ContribStats(ctx, ghClient)
	if err != nil {
		return err
	}

	fmt.Println("total commits:", numCommits)
	return nil
}
