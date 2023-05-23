package github

import (
	"time"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/shurcooL/githubv4"
)

func (c *AuthenticatedGitHubContext) NumCommits() (int, error) {
	var repoQuery struct {
		Viewer struct {
			ContributionsCollection struct {
				CommitContributionsByRepository []struct {
					Repository struct {
						NameWithOwner    string
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
				}
			} `graphql:"contributionsCollection(from: $from)"`
		}
	}

	vars := map[string]interface{}{
		"author_id": c.viewer.Viewer.ID,
		"since":     githubv4.GitTimestamp{Time: c.viewer.Viewer.CreatedAt.Time},
	}

	perRepo := map[string]int{}
	for from := c.viewer.Viewer.CreatedAt.Time; !from.After(time.Now()); from = from.AddDate(1, 0, 0) {
		vars["from"] = githubv4.DateTime{Time: from}
		err := c.githubClient.Query(c.ctx, &repoQuery, vars)
		if err != nil {
			return 0, err
		}

		total := 0
		for _, repo := range repoQuery.Viewer.ContributionsCollection.CommitContributionsByRepository {
			cnt := repo.Repository.DefaultBranchRef.Target.SpreadCommits.History.TotalCount
			total += cnt
			perRepo[repo.Repository.NameWithOwner] = cnt
		}
		log.Logvf("%d commits in %d repos, since %v\n", total, len(repoQuery.Viewer.ContributionsCollection.CommitContributionsByRepository), from)
	}

	totalCnt := 0
	log.Logvf("\nTotal commits per repo:\n")
	for name, cnt := range perRepo {
		log.Logvf("  %s%s: %d\n", icons.Repo, name, cnt)
		totalCnt += cnt
	}
	log.Logvf("\n")

	return totalCnt, nil
}
