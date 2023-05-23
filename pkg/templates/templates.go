package templates

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/liouk/gh-stats/pkg/stats"
)

type templateContainer struct {
	NumRepos   int
	NumForks   int
	NumPulls   int
	NumCommits int
	NumReviews int
	Languages  []*stats.Lang
	Extras     map[string]string
}

func Render(file, outfile string, stats *stats.GitHubViewerStats, extras map[string]string) error {
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
		NumRepos:   stats.RepoStats.NumRepos,
		NumForks:   stats.RepoStats.NumForks,
		NumPulls:   stats.RepoStats.NumPulls,
		NumCommits: stats.CommitStats.NumCommits,
		NumReviews: stats.ReviewStats.NumReviews,
		Languages:  stats.LangStats.Languages,
		Extras:     extras,
	}

	if err := tmpl.Execute(f, values); err != nil {
		return err
	}

	return nil
}
