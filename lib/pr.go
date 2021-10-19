package lib

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
)

func CreatePullRequest(ctx context.Context, client *github.Client, yes bool, owner, repo, from, to, title, body string) (string, error) {
	pull := github.NewPullRequest{
		Title: &title,
		Body:  &body,
		Base:  &to,
		Head:  &from,
	}

	isCreatePR := yes

	if !isCreatePR {
		fmt.Println("Do you want to create a pull request? [y/n]")
		var answer string
		fmt.Scanln(&answer)
		if answer == "y" {
			isCreatePR = true
		}
	}

	if isCreatePR {
		pr, _, err := client.PullRequests.Create(ctx, owner, repo, &pull)
		if err != nil {
			return "", fmt.Errorf("CreatePullRequest github.Client.PullRequests.Create: %v", err)
		}

		return pr.GetHTMLURL(), nil
	} else {
		return "", nil
	}
}
