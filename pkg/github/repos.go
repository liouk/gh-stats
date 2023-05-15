package github

func (c *AuthenticatedGitHubContext) NumRepos() (int, error) {
	var repos struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
			} `graphql:"repositories(isFork: false, privacy: PUBLIC, ownerAffiliations: OWNER)"`
		}
	}
	if err := c.githubClient.Query(c.ctx, &repos, nil); err != nil {
		return 0, err
	}

	return repos.Viewer.Repositories.TotalCount, nil
}

func (c *AuthenticatedGitHubContext) NumForks() (int, error) {
	var forks struct {
		Viewer struct {
			Repositories struct {
				TotalCount int
			} `graphql:"repositories(isFork: true, privacy: PUBLIC)"`
		}
	}
	if err := c.githubClient.Query(c.ctx, &forks, nil); err != nil {
		return 0, err
	}

	return forks.Viewer.Repositories.TotalCount, nil
}
