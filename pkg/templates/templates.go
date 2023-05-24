package templates

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/liouk/gh-stats/pkg/stats"
)

type templateContainer struct {
	NumRepos       int
	NumForks       int
	NumPulls       int
	NumOpenPulls   int
	NumClosedPulls int
	NumMergedPulls int
	NumCommits     int
	NumReviews     int
	Languages      []*stats.Lang
	Extras         map[string]interface{}
	User           string
}

func Render(file, outfile, githubUsername string, stats *stats.GitHubViewerStats, extras map[string]interface{}) error {
	tmplName := filepath.Base(file)
	tmpl, err := template.New(tmplName).ParseFiles(file)
	if err != nil {
		return err
	}

	f, err := os.Create(outfile)
	if err != nil {
		return err
	}

	values := templateContainer{
		User:           githubUsername,
		NumRepos:       stats.RepoStats.NumRepos,
		NumForks:       stats.RepoStats.NumForks,
		NumPulls:       stats.PullStats.TotalCount,
		NumOpenPulls:   stats.PullStats.OpenCount,
		NumClosedPulls: stats.PullStats.ClosedCount,
		NumMergedPulls: stats.PullStats.MergedCount,
		NumCommits:     stats.CommitStats.NumCommits,
		NumReviews:     stats.ReviewStats.NumReviews,
		Languages:      stats.LangStats.Languages,
		Extras:         extras,
	}

	if err := tmpl.Execute(f, values); err != nil {
		return err
	}

	return nil
}
