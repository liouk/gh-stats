package cmd

import "github.com/urfave/cli"

func NewCLIApp() *cli.App {
	return &cli.App{
		Name:   "gh-stats",
		Usage:  "Generate GitHub user stats",
		Action: root,
	}
}

func root(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}
