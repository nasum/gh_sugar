package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/v63/github"
)

func CountPR(ctx context.Context, client *github.Client, owner, targetName, repo, from, to string) (int, error) {
	fromTime, err := ParseDateFromString(from)
	if err != nil {
		return 0, err
	}

	toTime, err := ParseDateFromString(to)
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

	filteredPRs, err := filterPRs(allPRs, fromTime, toTime, targetName)
	if err != nil {
		return 0, err
	}

	return len(filteredPRs), nil
}

func fetchPRs(ctx context.Context, client *github.Client, owner, repo string) ([]*github.PullRequest, error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "fetching PRs... "
	s.FinalMSG = "fetching PRs done\n"
	s.Start()
	defer s.Stop()

	opts := &github.PullRequestListOptions{
		State: "all",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	var allPRs []*github.PullRequest
	for {
		s.Suffix = fmt.Sprintf(" page: %d", opts.Page)
		prs, resp, err := client.PullRequests.List(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		allPRs = append(allPRs, prs...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allPRs, nil
}

func filterPRs(allPRs []*github.PullRequest, fromTime, toTime time.Time, targetName string) ([]*github.PullRequest, error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "filtering PRs... "
	s.FinalMSG = "filtering PRs done\n"
	s.Start()
	defer s.Stop()

	var filteredPRs []*github.PullRequest
	for _, pr := range allPRs {
		if pr.CreatedAt != nil && pr.CreatedAt.After(fromTime) && pr.CreatedAt.Before(toTime) && *pr.User.Login == targetName {
			filteredPRs = append(filteredPRs, pr)
		}
	}

	return filteredPRs, nil
}
