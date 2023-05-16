package github

import (
	"fmt"

	"github.com/shurcooL/githubv4"
)

func (c *AuthenticatedGitHubContext) NumReviews() (int, error) {
	var search struct {
		Search struct {
			IssueCount int
		} `graphql:"search(type: ISSUE, query: $query)"`
	}

	vars := map[string]interface{}{
		"query": githubv4.String(fmt.Sprintf("is:public type:pr assignee:%s", c.viewer.Viewer.Login)),
	}

	if err := c.githubClient.Query(c.ctx, &search, vars); err != nil {
		return 0, err
	}

	return search.Search.IssueCount, nil
}
