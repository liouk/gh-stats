package github

import (
	"context"
	"time"

	"github.com/liouk/gh-stats/pkg/auth"
	"github.com/shurcooL/githubv4"
)

type viewerInfo struct {
	Viewer struct {
		Login     string
		CreatedAt time.Time
	}
}

type AuthenticatedGitHubContext struct {
	ctx          context.Context
	githubClient *githubv4.Client
	viewer       viewerInfo
}

func NewAuthenticatedGitHubContext() (*AuthenticatedGitHubContext, error) {
	ctx := &AuthenticatedGitHubContext{
		ctx:          context.Background(),
		githubClient: githubv4.NewClient(auth.NewOAuth2Client()),
	}

	if err := ctx.githubClient.Query(ctx.ctx, &ctx.viewer, nil); err != nil {
		return nil, err
	}

	return ctx, nil
}
