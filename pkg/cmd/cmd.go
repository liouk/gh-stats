package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/urfave/cli"
)

func NewCLIApp() *cli.App {
	return &cli.App{
		Name:   "gh-stats",
		Usage:  "Generate GitHub user stats",
		Action: root,
	}
}

func root(c *cli.Context) error {
	client := github.NewAuthenticatedClient()
	return github.RunSimpleQuery(client)
}
