package cmd

import (
	"github.com/liouk/gh-stats/pkg/stats"
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
	gh, err := initCmd(cCtx)
	if err != nil {
		return err
	}

	stats := &stats.GitHubViewerStats{LangStats: &stats.GitHubLangStats{}}
	stats.LangStats.Languages, err = gh.LangStats(cCtx.Int("num"), cCtx.StringSlice("ignore"))
	if err != nil {
		return err
	}

	return writeStats(cCtx, gh, stats)
}
