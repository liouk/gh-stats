package github

import (
	"sort"
	"strings"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
)

type set map[string]struct{}

type Lang struct {
	Name string
	Perc float32
}

func (c *AuthenticatedGitHubContext) LangStats(maxNum int, ignore []string) ([]*Lang, error) {
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
			} `graphql:"repositories(first: 100, ownerAffiliations: OWNER, isFork: false, privacy: PUBLIC)"`
		}
	}

	err := c.githubClient.Query(c.ctx, &langQuery, nil)
	if err != nil {
		return nil, err
	}

	ignoreSet := &set{}
	ignoreSet.AddAll(ignore)

	bytes := map[string]int{}
	totalBytes := 0
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

	lang := make([]*Lang, 0)
	for k, v := range bytes {
		lang = append(lang, &Lang{k, 100.0 * float32(v) / float32(totalBytes)})
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

func (s set) AddAll(slice []string) {
	for _, elem := range slice {
		s[strings.ToLower(elem)] = struct{}{}
	}
}

func (s set) ContainsIgnoreCase(elem string) bool {
	_, ok := s[strings.ToLower(elem)]
	return ok
}
