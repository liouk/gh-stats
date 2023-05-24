package github

import (
	"github.com/liouk/gh-stats/pkg/stats"
	"github.com/shurcooL/githubv4"
)

func (c *AuthenticatedGitHubContext) NumPulls() (*stats.GitHubPullStats, error) {
	var pulls struct {
		Viewer struct {
			PullRequests struct {
				TotalCount int
			} `graphql:"pullRequests(states: $states)"`
		}
	}

	cnt := []int{}
	total := 0
	vars := map[string]interface{}{}

	for _, state := range []githubv4.PullRequestState{
		githubv4.PullRequestStateOpen,
		githubv4.PullRequestStateClosed,
		githubv4.PullRequestStateMerged,
	} {
		vars["states"] = []githubv4.PullRequestState{state}
		if err := c.githubClient.Query(c.ctx, &pulls, vars); err != nil {
			return nil, err
		}

		total += pulls.Viewer.PullRequests.TotalCount
		cnt = append(cnt, pulls.Viewer.PullRequests.TotalCount)
	}

	return &stats.GitHubPullStats{
		TotalCount:  total,
		OpenCount:   cnt[0],
		ClosedCount: cnt[1],
		MergedCount: cnt[2],
	}, nil
}
