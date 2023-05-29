package github

import (
	"sort"
	"strings"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/shurcooL/githubv4"
)

type set map[string]struct{}

func (s set) AddAll(slice []string) {
	for _, elem := range slice {
		s[strings.ToLower(elem)] = struct{}{}
	}
}

func (s set) ContainsIgnoreCase(elem string) bool {
	_, ok := s[strings.ToLower(elem)]
	return ok
}

func (c *AuthenticatedGitHubContext) LangStats(maxNum int, ignore []string) ([]*stats.Lang, error) {
	var langQuery struct {
		Viewer struct {
			Repositories struct {
				Nodes []struct {
					NameWithOwner string
					Languages     struct {
						Edges []struct {
							Size int
							Node struct {
								Name string
							}
						}
					} `graphql:"languages(first: 100)"`
				}
				PageInfo struct {
					HasNextPage bool
					EndCursor   githubv4.String
				}
			} `graphql:"repositories(first: 100, after: $after, ownerAffiliations: OWNER, isFork: false, privacy: PUBLIC)"`
		}
	}

	ignoreSet := &set{}
	ignoreSet.AddAll(ignore)

	bytes := map[string]int{}
	totalBytes := 0

	vars := map[string]interface{}{
		"after": (*githubv4.String)(nil),
	}

	for {
		err := c.githubClient.Query(c.ctx, &langQuery, vars)
		if err != nil {
			return nil, err
		}

		for _, repo := range langQuery.Viewer.Repositories.Nodes {
			log.Logvf("%s%s\n", icons.Repo, repo.NameWithOwner)
			for _, lang := range repo.Languages.Edges {
				if ignoreSet.ContainsIgnoreCase(lang.Node.Name) {
					continue
				}

				n := lang.Node.Name
				if _, ok := bytes[n]; !ok {
					bytes[n] = 0
				}
				bytes[n] += lang.Size
				totalBytes += lang.Size
				log.Logvf("  %s: %d\n", n, lang.Size)
			}
		}

		if !langQuery.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}

		vars["after"] = langQuery.Viewer.Repositories.PageInfo.EndCursor
	}

	lang := make([]*stats.Lang, 0)
	for k, v := range bytes {
		lang = append(lang, &stats.Lang{Name: k, Perc: 100.0 * float32(v) / float32(totalBytes)})
	}

	lim := len(lang)
	if lim > maxNum {
		lim = maxNum
	}

	sort.Slice(lang, func(i, j int) bool {
		return lang[i].Perc > lang[j].Perc
	})
	return lang[0:lim], nil
}
