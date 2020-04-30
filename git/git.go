package git

import (
	"context"
	github "github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
	"k8s-enhancements/utils"
)

var gitClient *github.Client

func InitGit(token string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	gitClient = github.NewClient(tc)
}

func GetMilestone(repo, milestoneName string) *github.Milestone {
	if milestones, _, err := gitClient.Issues.ListMilestones(context.Background(), "kubernetes", repo, &github.MilestoneListOptions{}); err != nil {
		panic(err)
	} else {
		for _, milestone := range milestones {
			if milestone.GetTitle() == milestoneName {
				return milestone
			}
		}
	}
	return nil
}

func ListIssues(repo, state, assignee, sortOptions string, milestones, labels []string, maxSize int) {
	issuesToList := make([]*github.Issue, 0)
	milestoneMap := make(map[string]bool, 0)

	for _, m := range milestones {
		milestoneMap[m] = true
	}

	listOptions := &github.IssueListByRepoOptions{
		State:    state,
		Assignee: assignee,
		Labels:   labels,
		Sort:     sortOptions,
	}

	for {
		listOptions.PerPage = maxSize
		if issues, resp, err := gitClient.Issues.ListByRepo(context.Background(), "kubernetes", repo, listOptions); err != nil {
			panic(err)
		} else {
			for _, issue := range issues {
				if !issue.IsPullRequest() {
					if len(milestoneMap) > 0 {
						if _, ok := milestoneMap[issue.GetMilestone().GetTitle()]; ok {
							issuesToList = append(issuesToList, issue)
						}
					} else {
						issuesToList = append(issuesToList, issue)
					}
				}
			}
			if resp.NextPage == 0 {
				break
			} else {
				listOptions.Page = resp.NextPage
			}
		}
	}
	utils.DisplayIssues(issuesToList)
}

func AddComment(owner, repo, comment string, issueID int)  {
	cmt := &github.IssueComment{Body: github.String(comment)}

	if _, _, err := gitClient.Issues.CreateComment(context.Background(), owner, repo, issueID, cmt); err != nil {
		panic(err)
	}
}