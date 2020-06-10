package utils

import (
	"github.com/olekukonko/tablewriter"
	"k8s-enhancements/models"
	"os"
	"strconv"
	"strings"
)

var tracker models.Tracker

func DisplaySummary(summary map[string]models.Summary) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Enhancement Status", "Stage", "Issue Count", "Issue ID", "Issue Description"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	var data [][]string

	for status, info := range summary {
		for stage, records := range info.IssueData {
			for _, record := range records {
				r := make([]string, 0)
				r = append(r, status)
				r = append(r, stage)
				r = append(r, strconv.Itoa(len(records)))
				r = append(r, record.IssueId)
				r = append(r, record.Title)
				data = append(data, r)
			}
		}
	}

	table.SetAutoMergeCells(true)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetColWidth(80)
	table.SetAutoWrapText(true)
	table.AppendBulk(data)
	table.Render()
}

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
	table.SetColWidth(52)
	table.SetColMinWidth(1, 50)
	table.SetColMinWidth(9, 10)
	table.SetAutoWrapText(true)
	table.AppendBulk(data)
	table.Render()
}

func transformTrackingStatus(status string) string {
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

func init() {
	tracker = GetTrackingData()
}
