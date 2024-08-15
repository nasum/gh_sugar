package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/v63/github"
)

func CountReview(ctx context.Context, client *github.Client, owner, targetName, repo, from, to string) (int, error) {
	fromTime, err := parseDateFromString(from)
	if err != nil {
		return 0, err
	}

	toTime, err := parseDateFromString(to)
	if err != nil {
		return 0, err
	}

	fmt.Println("owner: ", owner)
	fmt.Println("repo: ", repo)
	fmt.Println("from: ", fromTime)
	fmt.Println("to: ", toTime)
	fmt.Println("target: ", targetName)

	allPRs, err := fetchPRs(ctx, client, owner, repo)
	if err != nil {
		return 0, err
	}

	filterdPRs, err := filterReviewPRs(allPRs, fromTime, toTime, targetName)
	if err != nil {
		return 0, err
	}

	return len(filterdPRs), nil
}

func filterReviewPRs(prs []*github.PullRequest, fromTime, toTime time.Time, targetName string) ([]*github.PullRequest, error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "filtering reviews... "
	s.FinalMSG = "filtering reviews done\n"
	s.Start()
	defer s.Stop()

	var filteredPRs []*github.PullRequest
	for _, pr := range prs {
		if pr.CreatedAt.After(fromTime) && pr.CreatedAt.Before(toTime) && pr.User.GetLogin() == targetName {
			filteredPRs = append(filteredPRs, pr)
		}
	}
	return filteredPRs, nil
}
