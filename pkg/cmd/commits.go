package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
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
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	return cmdCommitsWithGitHubContext(cCtx, gh)
}

func cmdCommitsWithGitHubContext(cCtx *cli.Context, gh *github.AuthenticatedGitHubContext) error {
	numCommits, err := gh.NumCommits(cCtx.StringSlice("in-depth"), cCtx.Bool("verbose"))
	if err != nil {
		return err
	}

	log.Logf("%s Commits: %d\n", icons.Commit, numCommits)
	return nil
}
