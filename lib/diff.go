package lib

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v35/github"
)

type PullRequest struct {
	Number int64
	Title  string
	Body   string
	URL    string
}

func (pr *PullRequest) ToString() string {
	return fmt.Sprintf("#%v %v", pr.Number, pr.Title)
}

func BranchDiff(ctx context.Context, client *github.Client, owner, repo, from, to string) (string, string, error) {
	comp, _, err := client.Repositories.CompareCommits(ctx, owner, repo, to, from)

	if err != nil {
		return "", "", fmt.Errorf("BranchDiff github.Client.Repositories.CompareCommits: %v", err)
	}

	re, err := regexp.Compile("#[0-9]+")

	if err != nil {
		return "", "", fmt.Errorf("BranchDiff regexp.Compile: %v", err)
	}

	var prNumberStArray []string

	for _, v := range comp.Commits {
		prNumberSt := re.FindString(*v.Commit.Message)
		if prNumberSt != "" {
			prNumberStArray = append(prNumberStArray, prNumberSt)
		}
	}

	var prArray []PullRequest

	for _, v := range prNumberStArray {
		prNum, err := strconv.Atoi(v[1:])

		if err != nil {
			return "", "", fmt.Errorf("BranchDiff strconv.Atoi: %v", err)
		}

		pr, _, err := client.PullRequests.Get(ctx, owner, repo, prNum)

		if err != nil {
			return "", "", fmt.Errorf("BranchDiff github.Client.PullRequests.Get: %v", err)
		}

		prStruct := PullRequest{
			Number: int64(prNum),
			Title:  pr.GetTitle(),
			Body:   pr.GetBody(),
			URL:    pr.GetURL(),
		}

		prArray = append(prArray, prStruct)
	}

	return createTitle(from, to), createBody(prArray), nil
}

func createTitle(from, to string) string {
	now := time.Now()
	nowSt := now.Format("2006/01/02 15:04:05")
	return fmt.Sprintf("%v %v to %v", nowSt, from, to)
}

func createBody(prArray []PullRequest) string {
	body := "# Diff\n"

	for _, v := range prArray {
		body = body + v.ToString() + "\n"
	}

	return body
}
