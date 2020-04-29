package utils

import (
	"github.com/google/go-github/v31/github"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

func DisplayIssues(issues []*github.Issue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Issue ID", "Issue Title", "Assignee", "Labels", "Milestone", "Last Updated", "PR Link"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	var data [][]string

	for _, issue := range issues {
		var labels []string
		var assignees []string
		for _, label := range issue.Labels {
			labels = append(labels, label.GetName())
		}

		for _, assignee := range issue.Assignees {
			assignees = append(assignees, assignee.GetLogin())
		}
		data = append(data, []string{strconv.Itoa(issue.GetNumber()), issue.GetTitle(), strings.Join(assignees, "\n"), strings.Join(labels, "\n"), issue.GetMilestone().GetTitle(), issue.GetUpdatedAt().Format("January 2, 2006"), issue.GetPullRequestLinks().GetURL()})
	}
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColMinWidth(1, 52)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()

}
