package github

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/liouk/gh-stats/pkg/auth"
	"github.com/liouk/gh-stats/pkg/icons"
	"github.com/liouk/gh-stats/pkg/log"
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
	logViewer(&ctx.viewer)

	return ctx, nil
}

func logViewer(info *viewerInfo) {
	titleStr := fmt.Sprintf("logged in as")
	userStr := fmt.Sprintf("%s%s", icons.GitHub, info.Viewer.Login)
	maxLen := len(titleStr)
	if len(userStr) > maxLen {
		maxLen = len(userStr)
	}

	sep := strings.Repeat("~", maxLen+4)

	log.Logf("%s\n  %s\n  %s\n%s\n\n", sep, titleStr, userStr, sep)
}
