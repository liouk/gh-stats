package github

func (c *AuthenticatedGitHubContext) NumPulls() (int, error) {
	var pulls struct {
		Viewer struct {
			PullRequests struct {
				TotalCount int
			}
		}
	}
	if err := c.githubClient.Query(c.ctx, &pulls, nil); err != nil {
		return 0, err
	}

	return pulls.Viewer.PullRequests.TotalCount, nil
}
