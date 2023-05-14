package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func NumPulls(ctx context.Context, c *githubv4.Client) (int, error) {
	var pulls struct {
		Viewer struct {
			PullRequests struct {
				TotalCount int
			}
		}
	}
	if err := c.Query(ctx, &pulls, nil); err != nil {
		return 0, err
	}

	return pulls.Viewer.PullRequests.TotalCount, nil
}
