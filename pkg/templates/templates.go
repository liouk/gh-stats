package templates

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/liouk/gh-stats/pkg/stats"
)

func Render(file, outfile string, stats *stats.GitHubViewerStats) error {
	tmplName := filepath.Base(file)
	tmpl, err := template.New(tmplName).ParseFiles(file)
	if err != nil {
		return err
	}

	f, err := os.Create(outfile)
	if err != nil {
		return err
	}

	return tmpl.Execute(f, stats)
}
