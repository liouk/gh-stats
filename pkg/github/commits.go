package github

import (
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/shurcooL/githubv4"
)

func (c *AuthenticatedGitHubContext) NumCommits() (int, error) {
	var repoQuery struct {
		Viewer struct {
			Repositories struct {
				Nodes []struct {
					Name             string
					DefaultBranchRef struct {
						Target struct {
							SpreadCommits struct {
								History struct {
									TotalCount int
								} `graphql:"history(since: $since, author: {id: $author_id})"`
							} `graphql:"... on Commit"`
						}
					}
				}
			} `graphql:"repositories(first: 100, affiliations: OWNER)"`
		}
	}

	vars := map[string]interface{}{
		"author_id": c.viewer.Viewer.ID,
		"since":     githubv4.GitTimestamp{Time: c.viewer.Viewer.CreatedAt.Time},
	}

	err := c.githubClient.Query(c.ctx, &repoQuery, vars)
	if err != nil {
		return 0, err
	}

	totalCnt := 0
	log.Logvf("Commits on default branch, per repo:\n")
	for _, repo := range repoQuery.Viewer.Repositories.Nodes {
		cnt := repo.DefaultBranchRef.Target.SpreadCommits.History.TotalCount
		log.Logvf("  %s%s: %d\n", icons.Repo, repo.Name, cnt)
		totalCnt += cnt
	}
	log.Logvf("\n")

	return totalCnt, nil
}
