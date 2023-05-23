package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/liouk/gh-stats/pkg/output"
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/liouk/gh-stats/pkg/templates"
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

	template := cCtx.String("template")
	if err := validateTemplateFlagValue(template); err != nil {
		return nil, err
	}

	if strings.EqualFold(outputType, "stdout") && template == "" {
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
			Usage:   "choose output type (ignored if --template is also present); values: stdout, json",
			Value:   "stdout",
		},
		&cli.StringFlag{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "render a template with stats (takes precedence over --output)",
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

func validateTemplateFlagValue(value string) error {
	_, err := os.Stat(value)
	return err
}

func writeStats(cCtx *cli.Context, stats *stats.GitHubViewerStats) error {
	var err error
	templateFile := cCtx.String("template")
	if templateFile != "" {
		err = templates.Render(templateFile, stats)
	} else {
		err = output.Print(os.Stdout, stats, cCtx.String("output"))
	}

	return err
}
