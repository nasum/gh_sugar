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

	prCmd := flag.NewFlagSet("pr", flag.ExitOnError)
	prOwner := prCmd.String("owner", "", "owner name")
	prRepo := prCmd.String("repo", "", "repository")
	prFrom := prCmd.String("from", "", "from branch")
	prTo := prCmd.String("to", "", "to branch")

	diffCmd := flag.NewFlagSet("diff", flag.ExitOnError)
	diffOwner := diffCmd.String("owner", "", "owner name")
	diffRepo := diffCmd.String("repo", "", "repository")
	diffFrom := diffCmd.String("from", "", "from branch")
	diffTo := diffCmd.String("to", "", "to branch")

	if len(os.Args) < 2 {
		fmt.Println("help")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "pr":
		prCmd.Parse(os.Args[2:])

		title, body, err := lib.BranchDiff(ctx, client, *prOwner, *prRepo, *prFrom, *prTo)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(-1)
		}

		prUrl, err := lib.CreatePullRequest(ctx, client, *prOwner, *prRepo, *prFrom, *prTo, title, body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(-1)
		}

		fmt.Printf("PullRequest Url: %v", prUrl)

	case "diff":
		diffCmd.Parse(os.Args[2:])
		title, body, err := lib.BranchDiff(ctx, client, *diffOwner, *diffRepo, *diffFrom, *diffTo)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(-1)
		}

		fmt.Println(title)
		fmt.Println(body)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
