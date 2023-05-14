package github

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
)

type contribQuery struct {
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

func ContribStats(ctx context.Context, c *githubv4.Client) (int, error) {
	var viewer struct {
		Viewer struct {
			CreatedAt time.Time
		}
	}
	err := c.Query(ctx, &viewer, nil)
	if err != nil {
		return 0, err
	}

	totalContrib := 0
	for fromTime := viewer.Viewer.CreatedAt; fromTime.Before(time.Now()); fromTime = fromTime.AddDate(1, 0, 0) {
		fmt.Println("FROM TIME:", fromTime)
		var contrib contribQuery
		vars := map[string]interface{}{
			"fromTime": githubv4.DateTime{Time: fromTime},
		}

		err = c.Query(ctx, &contrib, vars)
		if err != nil {
			return 0, err
		}

		for _, repo := range contrib.Viewer.ContributionsCollection.CommitContributionsByRepository {
			fmt.Println(repo.Repository.Name, repo.Contributions.TotalCount)
			totalContrib += repo.Contributions.TotalCount
		}

	}

	return totalContrib, nil
}
