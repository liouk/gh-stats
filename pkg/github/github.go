package github

import (
	"github.com/liouk/gh-stats/pkg/auth"
	"github.com/shurcooL/githubv4"
)

func NewAuthenticatedClient() *githubv4.Client {
	return githubv4.NewClient(auth.NewOAuth2Client())
}
