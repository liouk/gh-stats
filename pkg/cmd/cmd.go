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
			newAllCmd(),
			newReposCmd(),
			newCommitsCmd(),
			newReviewsCmd(),
			newLangCmd(),
		},
	}
}

// to be used in each command to avoid inconvenient urfave/cli positioning
func flags(flags ...cli.Flag) []cli.Flag {
	baseFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "display verbose information",
		},
	}

	return append(baseFlags, flags...)
}
