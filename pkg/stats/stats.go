package stats

type GitHubViewerStats struct {
	RepoStats   *GitHubRepoStats   `json:"repoStats"`
	CommitStats *GitHubCommitStats `json:"commitStats"`
	ReviewStats *GitHubReviewStats `json:"reviewStats"`
	LangStats   *GitHubLangStats   `json:"languageStats"`
}

type GitHubRepoStats struct {
	NumRepos int `json:"reposTotalCount"`
	NumForks int `json:"forksTotalCount"`
	NumPulls int `json:"pullsTotalCount"`
}

type GitHubCommitStats struct {
	NumCommits int `json:"totalCount"`
}

type GitHubReviewStats struct {
	NumReviews int `json:"totalCount"`
}

type GitHubLangStats struct {
	Languages []*Lang `json:"languages"`
}

type Lang struct {
	Name string  `json:"name"`
	Perc float32 `json:"percent"`
}
