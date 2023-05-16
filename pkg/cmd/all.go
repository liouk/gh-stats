package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/urfave/cli/v2"
)

func newAllCmd() *cli.Command {
	return &cli.Command{
		Name:   "all",
		Usage:  "Gets all stats (repos, commits, reviews)",
		Action: cmdAll,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "in-depth",
				Usage: "list of repo names that will be analysed in-depth",
			},
		},
	}
}

func cmdAll(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	if err := cmdReposWithGitHubContext(cCtx, gh); err != nil {
		return err
	}

	if err := cmdCommitsWithGitHubContext(cCtx, gh); err != nil {
		return err
	}

	return nil
}
