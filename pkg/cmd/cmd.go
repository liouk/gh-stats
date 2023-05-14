package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewCLIApp() *cli.App {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Shows help",
	}

	return &cli.App{
		Name:  "gh-stats",
		Usage: "Generate GitHub user stats",
		Commands: []*cli.Command{
			newReposCmd(),
			newContribCmd(),
		},
	}
}
