package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/nasum/gh_sugar/lib"
)

func main() {
	ctx := context.Background()
	client := lib.NewClient(ctx)

	if len(os.Args) < 2 {
		fmt.Println("help")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "pr":
		cmd := flag.NewFlagSet("pr", flag.ExitOnError)
		owner := cmd.String("owner", "", "owner name")
		repo := cmd.String("repo", "", "repository")
		from := cmd.String("from", "", "from branch")
		to := cmd.String("to", "", "to branch")
		yes := cmd.Bool("yes", false, "yes")
		cmd.Parse(os.Args[2:])

		title, body, err := lib.BranchDiff(ctx, client, *owner, *repo, *from, *to)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		fmt.Println(title)
		fmt.Println(body)

		url, err := lib.CreatePullRequest(ctx, client, *yes, *owner, *repo, *from, *to, title, body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		if url != "" {
			fmt.Printf("PullRequest Url: %v\n", url)
		}
	case "diff":
		cmd := flag.NewFlagSet("diff", flag.ExitOnError)
		owner := cmd.String("owner", "", "owner name")
		repo := cmd.String("repo", "", "repository")
		from := cmd.String("from", "", "from branch")
		to := cmd.String("to", "", "to branch")
		cmd.Parse(os.Args[2:])
		title, body, err := lib.BranchDiff(ctx, client, *owner, *repo, *from, *to)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}

		fmt.Println(title)
		fmt.Println(body)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
