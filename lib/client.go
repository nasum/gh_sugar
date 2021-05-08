package lib

import (
	"context"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

// create github.Client
func NewClient(ctx context.Context) *github.Client {
	githubAccessToken := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
