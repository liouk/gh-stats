package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func NumRepos(ctx context.Context, c *githubv4.Client) (int, error) {
	var repos struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
			} `graphql:"repositories(isFork: false, privacy: PUBLIC, ownerAffiliations: OWNER)"`
		}
	}
	if err := c.Query(ctx, &repos, nil); err != nil {
		return 0, err
	}

	return repos.Viewer.Repositories.TotalCount, nil
}

func NumForks(ctx context.Context, c *githubv4.Client) (int, error) {
	var forks struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
			} `graphql:"repositories(isFork: true, privacy: PUBLIC)"`
		}
	}
	if err := c.Query(ctx, &forks, nil); err != nil {
		return 0, err
	}

	return forks.Viewer.Repositories.TotalCount, nil
}
