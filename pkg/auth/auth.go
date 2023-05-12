package auth

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const ghTokenEnvVar = "GITHUB_TOKEN"

func NewOAuth2Client() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(ghTokenEnvVar)},
	)

	return oauth2.NewClient(context.Background(), src)
}
