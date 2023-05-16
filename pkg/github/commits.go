package github

import (
	"fmt"
	"time"

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

func (c *AuthenticatedGitHubContext) NumCommits(inDepth []string) (int, error) {
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
	for fromTime := viewer.Viewer.CreatedAt; fromTime.Before(time.Now()); fromTime = fromTime.AddDate(1, 0, 0) {
		fmt.Println("FROM TIME:", fromTime)
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
				fmt.Printf("(skipping %s; will analyse in-depth)", repo.Repository.Name)
				continue
			}

			fmt.Println(repo.Repository.Name, repo.Contributions.TotalCount)
			totalCommits += repo.Contributions.TotalCount
		}

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
