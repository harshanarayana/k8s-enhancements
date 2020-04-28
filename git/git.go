package git

import (
	"context"
	"fmt"
	github "github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
)

var gitClient *github.Client

func InitGit(token string)  {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	gitClient = github.NewClient(tc)
}

func ListIssues(milestone, state, assignee string, labels []string, maxSize int)  {
	if issues, _, err := gitClient.Issues.ListByRepo(context.Background(), "kubernetes", "enhancements", &github.IssueListByRepoOptions{
		Milestone: milestone,
		State: state,
		Assignee: assignee,
		Labels: labels,
		ListOptions: github.ListOptions{
			PerPage: maxSize,
		},
	}); err != nil {
		panic(err)
	} else {
		for _, issue := range issues {
			if ! issue.IsPullRequest() {
				fmt.Println(issue.GetTitle())
			}
		}
	}
}