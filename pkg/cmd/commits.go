package cmd

import (
	"fmt"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli/v2"
)

func newCommitsCmd() *cli.Command {
	return &cli.Command{
		Name:   "commits",
		Usage:  "Gets commits stats",
		Action: cmdCommits,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "in-depth",
				Usage: "list of repo names that will be analysed in-depth",
			},
		},
	}
}

func cmdCommits(cCtx *cli.Context) error {
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	numCommits, err := gh.NumCommits(cCtx.StringSlice("in-depth"))
	if err != nil {
		return err
	}

	fmt.Println("total commits:", numCommits)
	return nil
}
