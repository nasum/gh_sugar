package lib

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
)

func CreatePullRequest(ctx context.Context, client *github.Client, owner, repo, from, to, title, body string) (string, error) {
	pull := github.NewPullRequest{
		Title: &title,
		Body:  &body,
		Base:  &to,
		Head:  &from,
	}
	pr, _, err := client.PullRequests.Create(ctx, owner, repo, &pull)

	if err != nil {
		return "", fmt.Errorf("CreatePullRequest github.Client.PullRequests.Create: %v", err)
	}

	return pr.GetHTMLURL(), nil
}
