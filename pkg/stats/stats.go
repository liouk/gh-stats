package stats

type GitHubViewerStats struct {
	RepoStats   *GitHubRepoStats
	CommitStats *GitHubCommitStats
	ReviewStats *GitHubReviewStats
	LangStats   *GitHubLangStats
}

type GitHubRepoStats struct {
	NumRepos int
	NumForks int
	NumPulls int
}

type GitHubCommitStats struct {
	NumCommits int
}

type GitHubReviewStats struct {
	NumReviews int
}

type GitHubLangStats struct {
	Languages []*Lang
}

type Lang struct {
	Name string
	Perc float32
}
