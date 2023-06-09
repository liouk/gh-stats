package app

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

	cmdLang = &cli.Command{
		Name:   "lang",
		Usage:  "Gets language stats",
		Action: runLang,
		Flags: flags(
			flagLangNum,
			flagLangIgnore,
		),
	}
)

func runLang(cCtx *cli.Context) error {
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
