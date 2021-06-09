package adapter

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func PullPayload() {
	client := github.NewClient(nil)

	newPR := &github.NewPullRequest{
		Title:               github.String("My awesome pull request"),
		Head:                github.String("branch_to_merge"),
		Base:                github.String("master"),
		Body:                github.String("This is the description of the PR created with the package `github.com/google/go-github/github`"),
		MaintainerCanModify: github.Bool(true),
	}

	pr, _, err := client.PullRequests.Create(context.Background(), "ZrH-Source",  , newPR)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
}
