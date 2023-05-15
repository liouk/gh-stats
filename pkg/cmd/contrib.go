package cmd

import (
	"fmt"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli/v2"
)

func newContribCmd() *cli.Command {
	return &cli.Command{
		Name:   "contrib",
		Usage:  "Gets contribution stats",
		Action: cmdContrib,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "in-depth",
				Usage: "list of repo names that will be analysed in-depth",
			},
		},
	}
}

func cmdContrib(cCtx *cli.Context) error {
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	numCommits, err := gh.ContribStats(cCtx.StringSlice("in-depth"))
	if err != nil {
		return err
	}

	fmt.Println("total commits:", numCommits)
	return nil
}
