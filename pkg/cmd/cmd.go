package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli/v2"
)

func NewCLIApp() *cli.App {
	ghClient := github.NewAuthenticatedClient()

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Shows help",
	}

	return &cli.App{
		Name:  "gh-stats",
		Usage: "Generate GitHub user stats",
		Commands: []*cli.Command{
			newReposCmd(ghClient),
			newContributionsCmd(ghClient),
		},
	}
}
