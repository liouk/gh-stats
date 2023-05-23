package cmd

import (
	"fmt"
	"strings"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/log"
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

func initCmd(cCtx *cli.Context) (*github.AuthenticatedGitHubContext, error) {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return nil, err
	}

	outputType := cCtx.String("output")
	if err := validateOutputFlagValue(outputType); err != nil {
		return nil, err
	}

	if strings.EqualFold(outputType, "stdout") {
		gh.LogViewer()
	}

	return gh, nil
}

// to be used in each command to avoid inconvenient urfave/cli positioning
func flags(flags ...cli.Flag) []cli.Flag {
	baseFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "display verbose information",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "choose output type; values: stdout, json",
			Value:   "stdout",
		},
	}

	return append(baseFlags, flags...)
}

func validateOutputFlagValue(value string) error {
	switch strings.ToLower(value) {
	case "stdout", "json":
		return nil
	}

	return fmt.Errorf("unsupported output type: %s", value)
}
