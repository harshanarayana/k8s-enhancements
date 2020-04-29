package utils

import (
	"github.com/olekukonko/tablewriter"
	"k8s-enhancements/models"
	"os"
	"strings"
)

var tracker models.Tracker

func DisplayRows(issues []*models.EnhancementRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Issue ID", "Issue Title", "Enhancement Status", "Stage Status", "Stage", "Sig", "KEP Owner", "Note", "Updated Date", "Tracking"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
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
		date, note, state := getTrackingState(issue.IssueID)
		data = append(
			data,
			[]string{
				issue.IssueID,
				issue.EnhancementTitle,
				issue.EnhancementStatus,
				issue.StageStatus,
				issue.Stage,
				issue.Sig,
				issue.KEPOwner,
				note,
				date,
				state,
			},
		)
	}
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColMinWidth(1, 75)
	table.SetColMinWidth(9, 10)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}

func transformTrackingStatus(status string) string  {
	switch strings.ToLower(status) {
	case "tracked", "track", "t":
		return "︎✔︎"
	case "ignored", "ignore", "i":
		return "✘"
	case "skip", "s":
		return "►"
	case "info":
		return "ℹ"
	case "watch", "monitor", "w", "m":
		return "◎"
	default:
		return "𖡄"
	}
}

func getTrackingState(issue string) (string, string, string) {
	if item, ok := tracker.Records[issue]; !ok {
		return "", "", transformTrackingStatus("unknown")
	} else {
		return item.Date, item.Note, transformTrackingStatus(item.Status)
	}
}

func init()  {
	tracker = GetTrackingData()
}