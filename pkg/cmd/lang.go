package cmd

import (
	"github.com/liouk/gh-stats/pkg/github"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/urfave/cli/v2"
)

var (
	flagLangNum = &cli.IntFlag{
		Name:    "num",
		Aliases: []string{"n"},
		Usage:   "number of languages to return stats for",
		Value:   5,
	}

	flagLangIgnore = &cli.StringSliceFlag{
		Name:    "ignore",
		Aliases: []string{"i"},
		Usage:   "list of languages to ignore (case-insensitive)",
	}
)

func newLangCmd() *cli.Command {
	return &cli.Command{
		Name:   "lang",
		Usage:  "Gets language stats",
		Action: cmdLang,
		Flags: flags(
			flagLangNum,
			flagLangIgnore,
		),
	}
}

func cmdLang(cCtx *cli.Context) error {
	log.Init(cCtx)
	gh, err := github.NewAuthenticatedGitHubContext()
	if err != nil {
		return err
	}

	return cmdLangWithGitHubContext(cCtx, gh)
}

func cmdLangWithGitHubContext(cCtx *cli.Context, gh *github.AuthenticatedGitHubContext) error {
	langStats, err := gh.LangStats(cCtx.Int("num"), cCtx.StringSlice("ignore"))
	if err != nil {
		return err
	}

	log.Logf("%sLanguage stats:\n", icons.Code)
	for _, lang := range langStats {
		log.Logf("  %s: %.2f%%\n", lang.Name, lang.Perc)
	}
	return nil
}
