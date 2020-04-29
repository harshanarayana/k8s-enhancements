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
	table.SetHeader([]string{"Issue ID", "Issue Title", "Assignee", "Labels", "Milestone", "Last Updated"})

	table.SetHeaderColor(
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
		for _, label := range issue.Labels {
			labels = append(labels, label.GetName())
		}
		data = append(data, []string{strconv.Itoa(issue.GetNumber()), issue.GetTitle(), issue.GetAssignee().GetLogin(), strings.Join(labels, "\n"), issue.GetMilestone().GetTitle(), issue.GetUpdatedAt().Format("January 2, 2006")})
	}
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColMinWidth(1, 75)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()

}
