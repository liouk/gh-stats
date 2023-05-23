package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/liouk/gh-stats/pkg/auth"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
	"github.com/shurcooL/githubv4"
)

type viewerInfo struct {
	Viewer struct {
		ID        githubv4.ID
		Login     githubv4.String
		CreatedAt githubv4.DateTime
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

func (gh *AuthenticatedGitHubContext) ViewerUsername() string {
	return string(gh.viewer.Viewer.Login)
}

func (gh *AuthenticatedGitHubContext) LogViewer() {
	titleStr := fmt.Sprintf("logged in as")
	userStr := fmt.Sprintf("%s%s", icons.GitHub, gh.viewer.Viewer.Login)
	maxLen := len(titleStr)
	if len(userStr) > maxLen {
		maxLen = len(userStr)
	}

	sep := strings.Repeat("~", maxLen+4)

	log.Logf("%s\n  %s\n  %s\n%s\n\n", sep, titleStr, userStr, sep)
}
