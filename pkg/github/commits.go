package github

import (
	"time"

	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/shurcooL/githubv4"
)

type commitsQuery struct {
	Viewer struct {
		ContributionsCollection struct {
			CommitContributionsByRepository []struct {
				Contributions struct {
					TotalCount int
					Nodes      []struct {
						CommitCount int
						OccurredAt  time.Time
					}
				} `graphql:"contributions(first:100)"`
				Repository struct {
					Name string
				}
			} `graphql:"commitContributionsByRepository(maxRepositories: 100)"`
		} `graphql:"contributionsCollection(from: $fromTime)"`
	}
}

func (c *AuthenticatedGitHubContext) NumCommits(inDepth []string, verbose bool) (int, error) {
	var viewer struct {
		Viewer struct {
			CreatedAt time.Time
		}
	}
	err := c.githubClient.Query(c.ctx, &viewer, nil)
	if err != nil {
		return 0, err
	}

	totalCommits := 0
	inDepthRepos := sliceToMap(inDepth)
	perRepo := map[string]int{}
	for fromTime := viewer.Viewer.CreatedAt; fromTime.Before(time.Now()); fromTime = fromTime.AddDate(1, 0, 0) {
		log.Logvf("FROM TIME: %s\n", fromTime)
		var commits commitsQuery
		vars := map[string]interface{}{
			"fromTime": githubv4.DateTime{Time: fromTime},
		}

		err = c.githubClient.Query(c.ctx, &commits, vars)
		if err != nil {
			return 0, err
		}

		for _, repo := range commits.Viewer.ContributionsCollection.CommitContributionsByRepository {
			if _, exists := inDepthRepos[repo.Repository.Name]; exists {
				// skip repos that will be analysed in-depth
				log.Logvf("(skipping %s; will analyse in-depth)\n", repo.Repository.Name)
				continue
			}

			if _, exists := perRepo[repo.Repository.Name]; !exists {
				perRepo[repo.Repository.Name] = 0
			}
			perRepo[repo.Repository.Name] += repo.Contributions.TotalCount

			log.Logvf("  %s %s %d\n", icons.Commit, repo.Repository.Name, repo.Contributions.TotalCount)
			totalCommits += repo.Contributions.TotalCount
		}
	}

	if log.Verbose() {
		log.Logvf("\nCommits per repo:\n")
		for repo, cnt := range perRepo {
			log.Logvf("%s %s: %d\n", icons.Commit, repo, cnt)
		}
		log.Logvf("\n")
	}

	totalInDepth, err := c.inDepthStats(inDepth)
	if err != nil {
		return 0, nil
	}

	return totalCommits + totalInDepth, nil
}

func (c *AuthenticatedGitHubContext) inDepthStats(inDepth []string) (int, error) {
	return 0, nil
}

func sliceToMap(slice []string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, s := range slice {
		m[s] = struct{}{}
	}

	return m
}
